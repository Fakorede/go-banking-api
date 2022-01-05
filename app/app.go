package app

import (
	"fmt"
	"go-banking-api/domain"
	"go-banking-api/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	// init/resolve app dependencies
	customerHandlers := NewCustomerHandlers(services.NewDefaultCustomerService(domain.NewCustomerRepositoryDB()))

	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods(http.MethodGet)

	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to banking api!")
}
