package app

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) newAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())

		return
	}

	parsedCustomerId, err := strconv.ParseInt(customerId, 10, 64)

	if err != nil {
		writeResponse(w, r, http.StatusUnprocessableEntity, err.Error())

		return
	}

	request.CustomerId = int(parsedCustomerId)

	account, appError := h.service.NewAccount(request)

	if appError != nil {
		writeResponse(w, r, appError.Code, appError.AsMessage())

		return
	}

	writeResponse(w, r, http.StatusCreated, account)
}

func (h AccountHandlers) makeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	accountId := vars["account_id"]

	var request dto.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())

		return
	}

	pAccountId, err := strconv.ParseInt(accountId, 10, 64)

	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())

		return
	}

	pCustomerId, err := strconv.ParseInt(customerId, 10, 64)

	if err != nil {
		writeResponse(w, r, http.StatusBadRequest, err.Error())
	}

	request.AccountId = int(pAccountId)
	request.CustomerId = int(pCustomerId)

	response, appErr := h.service.MakeTransaction(request)

	if appErr != nil {
		writeResponse(w, r, appErr.Code, appErr.AsMessage())

		return
	}

	writeResponse(w, r, http.StatusCreated, response)
}

func NewAccountHandlers(service service.AccountService) AccountHandlers {
	return AccountHandlers{
		service: service,
	}
}
