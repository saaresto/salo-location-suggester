package search

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

func ConvertResponse(place SaloResponse) TPResponse {
	formattedPlace := TPResponse{}
	formattedPlace.Slug = place.Code
	formattedPlace.Title = place.Name
	switch place.PlaceType {
	case "airport":
		formattedPlace.Subtitle = place.CityName
	default:
		formattedPlace.Subtitle = place.CountryName
	}
	return formattedPlace
}
