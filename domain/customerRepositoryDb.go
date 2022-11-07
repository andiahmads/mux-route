package domain

import (
	"database/sql"
	"log"
	"mux-route/helper"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *helper.AppError) {

	var err error
	customers := make([]Customer, 0)
	if status == "" {
		sql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = d.client.Select(&customers, sql)
	} else {
		sql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, sql, status)
	}

	if err != nil {
		return nil, helper.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *helper.AppError) {
	findById := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	var c Customer
	err := d.client.Get(&c, findById, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.NewNotFoundError("customers not found")
		} else {
			log.Println("Error while scanning customers", err.Error())
			return nil, helper.NewUnexpectedError("enexpected database error")
		}
	}

	return &c, nil
}
