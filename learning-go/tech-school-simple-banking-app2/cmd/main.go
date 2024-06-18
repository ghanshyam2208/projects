package main

import (
	"banking_app2/api/handlers"
	"banking_app2/api/repositories"
	"banking_app2/api/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the repository
	customerRepo := repositories.NewCustomerRepository()

	// Initialize the service with the repository
	customerService := services.NewCustomerService(customerRepo)

	// Initialize the handler with the service
	customerHandler := handlers.CustomerHandler{Service: customerService}

	// Set up the router and define the routes
	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.Service.GetAllCustomers).Methods(http.MethodGet)

	// Start the server
	http.ListenAndServe(":8080", router)
}
