package app

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Customers struct {
	Name    string `json:"fullName"`
	City    string `json: "city"`
	Zipcode string `json: "zipCode"`
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, World!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customers{
		{
			Name:    "kiran",
			City:    "thane",
			Zipcode: "421503",
		},
		{
			Name:    "ghanshyam",
			City:    "thane",
			Zipcode: "421503",
		},
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}
