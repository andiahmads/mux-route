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

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.GetAllCustomer)

	log.Fatal(http.ListenAndServe("localhost:8082", router))

}
