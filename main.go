package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const API_URL string = "https://places.aviasales.ru/v2/places.json"

type TPResponse struct {
	Slug     string `json:"slug"`
	Subtitle string `json:"subtitle"`
	Title    string `json:"title"`
}

type SaloResponse struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	CountryName string `json:"country_name"`
	CityName    string `json:"city_name"`
	PlaceType   string `json:"type"`
}

func main() {
	log.Println("Starting suggester service")

	http.HandleFunc("/search", searchHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var (
		term   = r.URL.Query().Get("term")
		locale = r.URL.Query().Get("locale")
	)

	if term == "" {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("Invalid request params: %s", r.URL.String())
		return
	}

	places, err := sendPlacesRequest(term, locale)

	if err == nil {
		log.Printf("Transforming response from aviasales api")
		formattedPlaces := []TPResponse{}
		for _, place := range places {
			formattedPlace := TPResponse{}

			formattedPlace.Slug = place.Code
			formattedPlace.Title = place.Name
			switch place.PlaceType {
			case "airport":
				formattedPlace.Subtitle = place.CityName
			default:
				formattedPlace.Subtitle = place.CountryName
			}

			formattedPlaces = append(formattedPlaces, formattedPlace)
		}
		response, _ := json.Marshal(formattedPlaces)
		w.Header().Add("Content-Type", "application/json;charset=UTF-8")
		_, _ = w.Write(response)
	} else {
		log.Printf("Looking for cached places")

	}
}

func sendPlacesRequest(term, locale string) ([]SaloResponse, error) {
	request, err := http.NewRequest("GET", API_URL, nil)
	if err != nil {
		log.Println("Could not create http request")
		return nil, err
	} else {
		query := request.URL.Query()
		query.Add("term", term)
		query.Add("locale", locale)
		request.URL.RawQuery = query.Encode()

		log.Printf("Sending request to %s", request.URL.String())
		client := &http.Client{
			Timeout: 3 * time.Second,
		}
		response, err := client.Do(request)
		if err != nil {
			log.Println("Could not get response from aviasales api", err)
			return nil, err
		}
		defer response.Body.Close()

		responseBody, _ := ioutil.ReadAll(response.Body)
		var formattedResponse []SaloResponse
		err = json.Unmarshal(responseBody, &formattedResponse)

		return formattedResponse, err
	}
}
