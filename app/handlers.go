package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"go-banking-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	service services.CustomerService
}

func NewCustomerHandlers(service services.CustomerService) CustomerHandlers {
	return CustomerHandlers {
		service: service,
	}
}

func (ch CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func (ch CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(customer_id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprintf(w, err.Message)
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}

}

func createCustomer(w http.ResponseWriter, r *http.Request) {

}
