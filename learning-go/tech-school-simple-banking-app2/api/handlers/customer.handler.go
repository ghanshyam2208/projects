package handlers

import (
	"banking_app2/api/services"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type CustomerHandler struct {
	Service services.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ch.Service.GetAllCustomers()

	if err != nil {
		log.Println(err)
		errors.New("getting error in customer handler")
	}

	writeResponse(w, http.StatusOK, customers)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
