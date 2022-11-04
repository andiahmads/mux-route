package app

import (
	"encoding/xml"
	"mux-route/service"
	"net/http"
)

type Customer struct {
	Name    string `json:full_name" xml:"name"`
	City    string `json:full_city" xml:"city"`
	Zipcode string `json:full_zipcode" xml:"zipcode"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, _ := ch.service.GetAllCustomer()

	if r.Header.Get("Content-Type") == "application/json" {
		w.Header().Add("content-type", "application/json")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("content-type", "application/json")
		xml.NewEncoder(w).Encode(customers)
	}
}
