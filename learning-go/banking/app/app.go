package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func StartApp() {
	router := mux.NewRouter()
	router.HandleFunc("/greet", greetHandler).Methods(http.MethodGet)
	router.HandleFunc("/customers", getCustomersHandler).Methods(http.MethodGet)
	router.HandleFunc("/customersXML", getCustomersHandlerXML).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomersHandler).Methods(http.MethodPost)
	router.HandleFunc("/customers/{id}", getCustomersByIdHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:    ":3000",
		Handler: router, // Set the ServeMux as the handler for the HTTP server
	}

	fmt.Println("Starting the banking app...")
	// Start the HTTP server
	server.ListenAndServe()
}
