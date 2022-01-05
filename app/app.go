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
	router := mux.NewRouter()

	// init/resolve app dependencies
	customerHandlers := NewCustomerHandlers(services.NewDefaultCustomerService(domain.NewCustomerRepositoryStub()))
	// customerHandlers := CustomerHandlers{
	// 	service: services.NewDefaultCustomerService(domain.NewCustomerRepositoryStub()),
	// }

	router.HandleFunc("/", hello).Methods(http.MethodGet)

	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", getCustomer).Methods(http.MethodGet)

	log.Println("app is starting...")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to banking api!")
}
