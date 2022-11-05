package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"mux-route/service"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"full_name"`
	City    string `json:"city" `
	Zipcode string `json:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("content-type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}
func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		w.WriteHeader(err.Code)
		fmt.Fprint(w, err.Message)
	} else {
		w.Header().Add("content-type", "application/json")
		json.NewEncoder(w).Encode(customer)
	}
}
