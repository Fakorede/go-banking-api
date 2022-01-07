package app

import (
	"encoding/json"
	"encoding/xml"
	"go-banking-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service services.CustomerService
}

func NewCustomerHandlers(service services.CustomerService) CustomerHandlers {
	return CustomerHandlers{
		service: service,
	}
}

func (h CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("status")
	customers, err := h.service.GetAllCustomers(param)
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	} else {
		writeResponse(w, r, http.StatusOK, customers)
	}
}

func (h CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customer_id := vars["customer_id"]

	customer, err := h.service.GetCustomer(customer_id)
	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())
	} else {
		writeResponse(w, r, http.StatusOK, customer)
	}

}

func writeResponse(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		err := xml.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(code)
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			panic(err)
		}
	}
}
