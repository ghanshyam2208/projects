package app

import (
	"log"
	"net/http"

	"github.com/ghanshyam2208/banking/domain"
	"github.com/ghanshyam2208/banking/service"
	"github.com/gorilla/mux"
)

func Start() {

	router := mux.NewRouter()

	// wiring
	handler := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// defining routes
	// router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", handler.getAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{customer_id:[0-9]+}", handler.GetCustomerById).Methods("GET")

	// starting the router
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}

// func NewCustomerHandler() CustomerHandler {
// 	var ch CustomerHandler
// 	ch.service = services.NewCustomerService(repositories.NewCustomerRepository())
// 	return ch
// }
