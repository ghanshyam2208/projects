package app

import (
	"encoding/json"
	"net/http"

	"github.com/ghanshyam2208/banking/service"
	"github.com/gorilla/mux"
)

// type Customers struct {
// 	Name    string `json:"fullName"`
// 	City    string `json: "city"`
// 	Zipcode string `json: "zipCode"`
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello, World!")
// }

// func getAllCustomers(w http.ResponseWriter, r *http.Request) {
// 	customers := []Customers{
// 		{
// 			Name:    "kiran",
// 			City:    "thane",
// 			Zipcode: "421503",
// 		},
// 		{
// 			Name:    "ghanshyam",
// 			City:    "thane",
// 			Zipcode: "421503",
// 		},
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(customers)
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customers, _ := ch.service.GetCustomerById(id)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
