package domain

import (
	"database/sql"
	"github.com/Fakorede/go-banking-api/errs"
	"github.com/Fakorede/go-banking-api/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDB struct {
	Client *sqlx.DB
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{
		Client: dbClient,
	}
}

func (db CustomerRepositoryDB) FindAll(status string) ([]Customer, *errs.AppError) {
	// var rows *sql.Rows
	var err error
	customers := []Customer{}

	if status == "" {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers"
		err = db.Client.Select(&customers, findAllSql)
		// rows, err = db.Client.Query(findAllSql)
	} else {
		findAllSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE status = ?"
		err = db.Client.Select(&customers, findAllSql, status)
		// rows, err = db.Client.Query(findAllSql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning customers row: " + err.Error())
	// 	return nil, errs.NewUnexpectedError("unexpected server error")
	// }

	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customers row: " + err.Error())
	// 		return nil, errs.NewUnexpectedError("unexpected server error")
	// 	}
	// 	customers = append(customers, c)
	// }

	return customers, nil
}

func (db CustomerRepositoryDB) FindByID(id string) (*Customer, *errs.AppError) {
	findSql := "SELECT customer_id, name, city, zipcode, date_of_birth, status FROM customers WHERE customer_id = ?"

	var c Customer

	//row := db.Client.QueryRow(findSql, id)
	// err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateofBirth, &c.Status)

	err := db.Client.Get(&c, findSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			logger.Error("Error while scanning customer row: " + err.Error())
			return nil, errs.NewUnexpectedError("unexpected server error")
		}
	}

	return &c, nil
}
