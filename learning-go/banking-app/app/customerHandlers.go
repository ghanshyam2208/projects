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
	status := r.URL.Query().Get("status")
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}

	customers, err := ch.service.GetAllCustomer(status)
	// w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customers)
	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) GetCustomerById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customers, err := ch.service.GetCustomerById(id)
	if err != nil {
		// w.Header().Add("Content-Type", "application/json")
		// w.WriteHeader(err.Code)
		// json.NewEncoder(w).Encode(err.AsMessage())
		writeResponse(w, err.Code, err.AsMessage())

	} else {
		// w.Header().Add("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(customers)
		writeResponse(w, http.StatusOK, customers)
	}

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
