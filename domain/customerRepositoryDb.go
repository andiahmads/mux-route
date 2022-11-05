package domain

import (
	"database/sql"
	"log"
	"mux-route/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	sql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"

	rows, err := d.client.Query(sql)
	if err != nil {
		log.Println("Error while querying customers table: ", err.Error())
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customer table: ", err.Error())
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *helper.AppError) {
	findById := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id = ?"

	row := d.client.QueryRow(findById, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
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

func NewCustomerRepositoryStDb() CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:endi@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}
