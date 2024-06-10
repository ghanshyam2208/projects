package main

import (
	"log"
	"simple_banking_app/api"
)

func main() {
	server := api.NewServer()
	if err := server.Router.Start(":8081"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
