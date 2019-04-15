package main

import (
	"fmt"
)

const API_URL string = "https://places.aviasales.ru/v2/places.json"

func main() {
	fmt.Println("Starting suggester service")
}

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
