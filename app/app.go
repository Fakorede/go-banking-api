package app

import (
	"fmt"
	"go-banking-api/domain"
	"go-banking-api/services"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func checkVars() {
	if 
		os.Getenv("SERVER_ADDRESS") == "" || 
		os.Getenv("SERVER_PORT") == "" || 
		os.Getenv("DB_USER") == "" || 
		os.Getenv("DB_PASSWORD") == "" || 
		os.Getenv("DB_HOST") == "" || 
		os.Getenv("DB_PORT") == "" || 
		os.Getenv("DB_NAME") == "" {
		log.Fatal("environment variables not defined!!!")
	}
}

func Start() {
	checkVars()

	// init/resolve app dependencies
	customerHandlers := NewCustomerHandlers(services.NewDefaultCustomerService(domain.NewCustomerRepositoryDB()))

	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods(http.MethodGet)

	router.HandleFunc("/customers", customerHandlers.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", customerHandlers.getCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers", createCustomer).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to banking api!")
}
