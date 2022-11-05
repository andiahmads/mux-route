package app

import (
	"log"
	"mux-route/domain"
	"mux-route/service"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStDb())}

	router.HandleFunc("/customers", ch.GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8082", router))

}
