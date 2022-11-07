package app

import (
	"mux-route/helper"
	"mux-route/logger"
	"mux-route/service"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomer(status)
	if err != nil {
		helper.WriteResponse(w, http.StatusNotFound, err.ASMessage())
		return
	} else {
		helper.WriteResponse(w, http.StatusOK, customers)
		return
	}
}
func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		helper.WriteResponse(w, err.Code, err.ASMessage())
		logger.Error(err.Message)
		return
	} else {
		helper.WriteResponse(w, http.StatusOK, customer)
		return
	}
}
