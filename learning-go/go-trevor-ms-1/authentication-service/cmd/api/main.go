package main

import (
	// data "authentication/models"
	data "authentication/cmd/api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const webPort = "8002"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {
	log.Printf("Starting authentication service on port %s \n", webPort)

	// TODO connect to DB

	// set up config
	app := Config{}

	// define http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}
