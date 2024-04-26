package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

func greetHandler(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprint(responseWriter, "Hello world")
}

func getCustomersHandler(responseWriter http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Tom", "Mumbai", "421503"},
		{"Robbin", "Mumbai", "421503"},
	}

	responseWriter.Header().Add("Content-Type", "application/json")
	json.NewEncoder(responseWriter).Encode(customers)
}

func getCustomersHandlerXML(responseWriter http.ResponseWriter, request *http.Request) {
	customers := []Customer{
		{"Tom", "Mumbai", "421503"},
		{"Robbin", "Mumbai", "421503"},
	}

	responseWriter.Header().Add("Content-Type", "application/xml")
	xml.NewEncoder(responseWriter).Encode(customers)
}

func createCustomersHandler(responseWriter http.ResponseWriter, request *http.Request) {

	fmt.Fprint(responseWriter, "this is post request")

}

func getCustomersByIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fmt.Fprint(responseWriter, vars["id"])
}
