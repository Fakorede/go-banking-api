package app

import (
	"go-banking-api/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func routes(ch handlers.CustomerHandlers, ah handlers.AccountHandlers) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", hello).Methods(http.MethodGet)

	router.HandleFunc("/customers", ch.GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	return router
}