package search

import "testing"

func TestConvertResponse(t *testing.T) {
	// given airport
	salo := SaloResponse{
		PlaceType:   "airport",
		CityName:    "Moscow",
		CountryName: "Russia",
	}
	tp := ConvertResponse(salo)

	// expect
	if tp.Subtitle != salo.CityName {
		t.Errorf("Expected %s to be airport's subtitle, got %s", salo.CityName, tp.Subtitle)
	}

	// given city
	salo.PlaceType = "city"
	tp = ConvertResponse(salo)

	// expect
	if tp.Subtitle != salo.CountryName {
		t.Errorf("Expected %s to be city's subtitle, got %s", salo.CountryName, tp.Subtitle)
	}
}
