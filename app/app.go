package app

import (
	"fmt"
	"log"
	"mux-route/domain"
	"mux-route/service"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func sanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variable not difined....")
	}
}

func Start() {

	sanityCheck()

	dbClient := getDbClient()

	router := mux.NewRouter()

	// call respository
	customerRepository := domain.NewCustomerRepositoryDb(dbClient)
	accountRepository := domain.NewAccountRepositoryDb(dbClient)

	// call service
	customerService := service.NewCustomerService(customerRepository)
	accountService := service.NewAccountService(accountRepository)

	// customer handler
	ch := CustomerHandlers{customerService}

	// account handler
	ah := AccountHandler{accountService}

	app, _ := newrelic.NewApplication(
		newrelic.ConfigAppName("banking"),
		newrelic.ConfigLicense("ae0bc44af2ed7136477b2668027f21c7e172NRAL"),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/customers", ch.GetAllCustomer)).Methods(http.MethodGet)
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/customers/{customer_id:[0-9]+}", ch.GetCustomer)).Methods(http.MethodGet)

	// account handler
	router.HandleFunc(newrelic.WrapHandleFunc(app, "/customers/{customer_id:[0-9]+}/account", ah.NewAccount)).Methods(http.MethodPost)

	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router))
}

func getDbClient() *sqlx.DB {

	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)

	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client

}
