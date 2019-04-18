package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	log.Println("Starting suggester service")

	envPort := os.Getenv("SERVER_PORT")
	if len(envPort) == 0 {
		envPort = "3001"
	}
	port, err := strconv.Atoi(envPort)
	if err != nil {
		log.Fatal(err)
	}
	srv := Server{int(port)}

	srv.Start()
}
