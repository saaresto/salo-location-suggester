package search

import (
	"crypto/tls"
	"encoding/json"
	"github.com/saaresto/salo-location-suggester/cache"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const ApiUrl string = "https://places.aviasales.ru/v2/places.json"

type Handler struct {
	cache      *cache.Cache
	apiTimeout time.Duration
}

func NewSearchHandler() *Handler {
	return &Handler{
		cache:      cache.InitializeCache(),
		apiTimeout: 3 * time.Second,
	}
}

// this is just for testing purposes
func (sh *Handler) SetApiTimeout(timeout time.Duration) {
	log.Printf("Setting search api timeout to %d", timeout)
	sh.apiTimeout = timeout
}

func (sh *Handler) HandleSearch(w http.ResponseWriter, r *http.Request) {
	// get unique id to log request based on timestamp (since generating uuid requires additional dependency)
	requestId := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)[6:15]

	var (
		term   = r.URL.Query().Get("term")
		locale = r.URL.Query().Get("locale")
	)

	if term == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("%s - Invalid request params: %s", requestId, r.URL.String())
		return
	}

	// todo consider using channels to implement timeout
	places, err := sh.sendPlacesRequest(term, locale, requestId)

	if err == nil {
		log.Printf("%s - Transforming response from aviasales api", requestId)
		var formattedPlaces []TPResponse
		for _, place := range places {
			formattedPlaces = append(formattedPlaces, ConvertResponse(place))
		}
		response, _ := json.Marshal(formattedPlaces)

		// put them in the cache asynchronously
		go func() {
			log.Printf("%s - Updating %d places in cache for term %s", requestId, len(formattedPlaces), term)
			sh.cache.PutValue(term, response)
		}()

		writeJsonResponse(w, response)
	} else {
		log.Printf("%s - Error occured while sending request to aviasales api: %v", requestId, err)
		value, ok := sh.cache.GetValue(term)
		if ok {
			log.Printf("%s - Returning cached values", requestId)
			value := value.([]byte)
			writeJsonResponse(w, value)
		} else {
			log.Printf("%s - No suitable records found in cache for term '%s'", requestId, term)
			writeJsonResponse(w, []byte("[]"))
		}
	}
}

func (sh *Handler) sendPlacesRequest(term, locale, requestId string) ([]SaloResponse, error) {
	request, err := http.NewRequest("GET", ApiUrl, nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	query.Add("term", term)
	query.Add("locale", locale)
	request.URL.RawQuery = query.Encode()

	log.Printf("%s - Sending request to %s", requestId, request.URL.String())

	tlsParam := os.Getenv("USE_TLS")
	var useTls bool
	if len(tlsParam) == 0 {
		useTls = true // turn it on by default
	} else {
		useTls, _ = strconv.ParseBool(tlsParam)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: !useTls},
	}

	client := http.Client{
		Timeout:   sh.apiTimeout,
		Transport: tr,
	}

	response, err := client.Do(request)
	if err != nil {
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
