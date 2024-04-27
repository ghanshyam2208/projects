package app

import (
	"net/http"

	"github.com/gnbaviskar/bankingApp/service"
)

// type ICustomerService interface {
// 	findAllCustomers() ([]domain.Customer, error)
// }

// type DefaultCustomerService struct {
// 	customerRepo domain.ICustomerRepo
// }

// func (defaultCustomerService DefaultCustomerService) findAllCustomers() ([]domain.Customer, error) {
// 	return defaultCustomerService.customerRepo.FindAll()
// }

// func NewCustomerService(customerRepo domain.ICustomerRepo) DefaultCustomerService {
// 	return DefaultCustomerService{customerRepo: customerRepo}
// }

type Customer struct {
	Name    string `json:"fullName" xml:"fullName"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipCode" xml:"zipCode"`
}

type CustomerHandler struct {
	customerService service.ICustomerService
}

// func greetHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	fmt.Fprint(responseWriter, "Hello world")
// }

func (customerHandler *CustomerHandler) getAllCustomers(responseWriter http.ResponseWriter, request *http.Request) {
	// customers := []Customer{
	// 	{"Tom", "Mumbai", "421503"},
	// 	{"Robbin", "Mumbai", "421503"},
	// }

	customers, error := customerHandler.customerService.findAllCustomers()
	// responseWriter.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(responseWriter).Encode(customers)
}

// func getCustomersHandlerXML(responseWriter http.ResponseWriter, request *http.Request) {
// 	customers := []Customer{
// 		{"Tom", "Mumbai", "421503"},
// 		{"Robbin", "Mumbai", "421503"},
// 	}

// 	responseWriter.Header().Add("Content-Type", "application/xml")
// 	xml.NewEncoder(responseWriter).Encode(customers)
// }

// func createCustomersHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	fmt.Fprint(responseWriter, "this is post request")
// }

// func getCustomersByIdHandler(responseWriter http.ResponseWriter, request *http.Request) {
// 	vars := mux.Vars(request)
// 	fmt.Fprint(responseWriter, vars["id"])
// }
