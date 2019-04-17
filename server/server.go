package server

import (
	"fmt"
	"github.com/saaresto/salo-location-suggester/search"
	"log"
	"net/http"
)

type Server struct {
	Port int
}

func (s *Server) Start() {
	http.HandleFunc("/search", search.SearchHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}
