package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
	"github.com/oopchi/banking/service"
	"github.com/oopchi/banking/util"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func NewCustomerHandlers(service service.CustomerService) CustomerHandlers {
	return CustomerHandlers{service: service}
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	var err *errs.AppError
	var customers []dto.CustomerResponse

	status := queries.Get("status")

	if status == "" {
		customers, err = ch.service.GetAllCustomer(nil)
	} else {
		statusCode := util.StatusCode(status)
		customers, err = ch.service.GetAllCustomer(&statusCode)
	}

	if err != nil {
		writeResponse(w, r, err.Code, err.AsMessage())

		return
	}

	writeResponse(w, r, http.StatusOK, customers)

}

func (ch *CustomerHandlers) getCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)

	if err != nil {

		writeResponse(w, r, err.Code, err.AsMessage())

		return
	}

	writeResponse(w, r, http.StatusOK, customer)

}
