package domain

import (
	"go-banking-api/errs"
	"go-banking-api/logger"
	"strconv"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	Client *sqlx.DB
}

func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		Client: dbClient,
	}
}

func (db AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type, amount, status) VALUES (?, ?, ?, ?, ?)"

	result, err := db.Client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	a.AccountId = strconv.FormatInt(id, 10)

	return &a, nil
}
