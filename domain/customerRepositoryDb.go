package domain

import (
	"database/sql"
	"go-banking-api/errs"
	"go-banking-api/logger"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	Client *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/go_banking_api")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDB{
		Client: db,
	}
}

func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error

	if status == "" {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		rows, err = db.Client.Query(findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		rows, err = db.Client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	customers := []Customer{} // make([]Customer, 0)

	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scanning customers row: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected server error")
		}

		customers = append(customers, c)
	}

	return customers, nil
}

func (db CustomerRepositoryDB) FindByID(id string) (Customer, *errs.AppError) {
	findSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	row := db.Client.QueryRow(findSql, id)
	var c Customer

	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return Customer{}, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer row: " + err.Error())
			return Customer{}, errs.NewUnexpectedError("unexpected server error")
		}
	}

	return c, nil
}
