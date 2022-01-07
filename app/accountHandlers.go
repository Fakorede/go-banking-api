package app

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

func (h AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
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
