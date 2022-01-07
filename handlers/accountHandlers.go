package handlers

import (
	"encoding/json"
	"go-banking-api/dto"
	"go-banking-api/services"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service services.AccountService
}

func NewAccountHandlers(service services.AccountService) AccountHandlers {
	return AccountHandlers{
		service: service,
	}
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := h.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, r, appError.Code, appError.AsMessage())
			return
		}
		writeResponse(w, r, http.StatusCreated, account)
	}
}

func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	customerId := mux.Vars(r)["customer_id"]
	accountId := mux.Vars(r)["account_id"]

	var request dto.TransactionRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		request.AccountId = accountId

		transaction, appError := h.service.MakeTransaction(request)
		if appError != nil {
			writeResponse(w, r, appError.Code, appError.AsMessage())
			return
		}
		writeResponse(w, r, http.StatusCreated, transaction)
	}
}
