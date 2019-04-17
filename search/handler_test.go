package search

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandler_HandleSearchMissingRequiredParams(t *testing.T) {
	// given
	sh := NewSearchHandler()

	req, err := http.NewRequest("GET", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	// when
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.HandleSearch)

	handler.ServeHTTP(rr, req)

	// expect
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestSearchHandler_HandleSearch(t *testing.T) {
	// given
	sh := NewSearchHandler()

	req, err := http.NewRequest("GET", "/search", nil)
	if err != nil {
		t.Fatal(err)
	}

	query := req.URL.Query()
	query.Add("term", "MOW")
	query.Add("locale", "ru")
	req.URL.RawQuery = query.Encode()

	// when
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(sh.HandleSearch)

	handler.ServeHTTP(rr, req)

	// expect
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
