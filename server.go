package main

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
	searchHandler := search.NewSearchHandler()
	http.HandleFunc("/search", searchHandler.HandleSearch)

	log.Printf("Server started at port :%d", s.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.Port), nil))
}
