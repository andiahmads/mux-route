package app

import (
	"encoding/json"
	"mux-route/dto"
	"mux-route/helper"
	"mux-route/service"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandler struct {
	service service.AccountService
}

func (h *AccountHandler) NewAccount(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	customerId := vars["customer_id"]

	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		helper.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	request.CustomerId = customerId
	account, appErr := h.service.NewAccount(request)
	if appErr != nil {
		helper.WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	helper.WriteResponse(w, http.StatusOK, account)
}
