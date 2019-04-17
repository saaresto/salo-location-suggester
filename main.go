package main

import (
	"log"
)

func main() {
	log.Println("Starting suggester service")

	srv := Server{3001}

	srv.Start()
}
