package main

import (
	"github.com/saaresto/salo-location-suggester/server"
	"log"
)

func main() {
	log.Println("Starting suggester service")

	srv := server.Server{3001}

	srv.Start()
}
