package main

import (
	// data "authentication/models"
	data "authentication/cmd/api/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = "8002"

type Config struct {
	DB     *sql.DB
	Models data.Models
}

var counts int64

func main() {
	log.Printf("Starting authentication service on port %s \n", webPort)

	conn := connectToDB()

	if conn == nil {
		log.Panic("Could not connect to pg")
		return
	}

	// TODO connect to DB

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

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

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")

	for {
		connection, err := openDB(dsn)

		if err != nil {
			log.Println("Postgres is not ready yet......")
			counts++
		} else {
			log.Println("Connected to pg")
			return connection
		}

		if counts > 10 {
			log.Println(err)
			return nil
		}

		log.Println("Backing off for two seconds.....")
		time.Sleep(2 * time.Second)

		continue
	}
}
