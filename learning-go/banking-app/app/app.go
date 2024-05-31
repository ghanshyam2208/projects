package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// defining routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)

	// starting the routerer
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
