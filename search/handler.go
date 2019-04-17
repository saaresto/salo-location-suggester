package search

import (
	"encoding/json"
	"github.com/saaresto/salo-location-suggester/cache"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const API_URL string = "https://places.aviasales.ru/v2/places.json"

type SearchHandler struct {
	cache      *cache.Cache
	apiTimeout time.Duration
}

func NewSearchHandler() *SearchHandler {
	return &SearchHandler{
		cache:      cache.InitializeCache(),
		apiTimeout: 3 * time.Second,
	}
}

// this is just for testing purposes
func (sh *SearchHandler) SetApiTimeout(timeout time.Duration) {
	log.Printf("Setting search api timeout to %d", timeout)
	sh.apiTimeout = timeout
}

func (sh *SearchHandler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	var (
		term   = r.URL.Query().Get("term")
		locale = r.URL.Query().Get("locale")
	)

	if term == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Invalid request params: %s", r.URL.String())
		return
	}

	places, err := sh.sendPlacesRequest(term, locale)

	if err == nil {
		log.Printf("Transforming response from aviasales api")
		formattedPlaces := []TPResponse{}
		for _, place := range places {
			formattedPlaces = append(formattedPlaces, ConvertResponse(place))
		}
		response, _ := json.Marshal(formattedPlaces)

		// put them in the cache asynchronously
		go func() {
			log.Printf("Updating %d places in cache", len(formattedPlaces))
			sh.cache.PutValue(term, response)
		}()

		writeJsonResponse(w, response)
	} else {
		log.Printf("Looking for cached places")
		value, ok := sh.cache.GetValue(term)
		if ok {
			value := value.([]byte)
			writeJsonResponse(w, value)
		} else {
			log.Printf("No suitable records found in cache for term '%s'", term)
			writeJsonResponse(w, make([]byte, 0))
		}
	}
}

func (sh *SearchHandler) sendPlacesRequest(term, locale string) ([]SaloResponse, error) {
	request, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Println("Could not create http request")
		return nil, err
	}

	query := request.URL.Query()
	query.Add("term", term)
	query.Add("locale", locale)
	request.URL.RawQuery = query.Encode()

	log.Printf("Sending request to %s", request.URL.String())
	client := http.Client{
		Timeout: sh.apiTimeout,
	}
	response, err := client.Do(request)
	if err != nil {
		log.Println("Could not get response from aviasales api:", err)
		return nil, err
	}
	defer response.Body.Close()

	responseBody, _ := ioutil.ReadAll(response.Body)
	var formattedResponse []SaloResponse
	err = json.Unmarshal(responseBody, &formattedResponse)

	return formattedResponse, err
}

func writeJsonResponse(w http.ResponseWriter, response []byte) {
	w.Header().Add("Content-Type", "application/json;charset=UTF-8")
	_, err := w.Write(response)

	if err != nil {
		log.Println("Could not write response:", err)
	}
}
