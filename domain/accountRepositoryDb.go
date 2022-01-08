package domain

import (
	"github.com/Fakorede/go-banking-api/errs"
	"github.com/Fakorede/go-banking-api/logger"
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

/**
 * make an entry in the transaction table + update the balance in the accounts table
 */
func (db AccountRepositoryDB) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	// start db transaction
	tx, err := db.Client.Begin()
	if err != nil {
		logger.Error("Error while starting new db transactin for bank account transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected server error")
	}

	// insert account transaction
	result, _ := tx.Exec(`INSERT INTO transactions(account_id, amount, transaction_type, transaction_date) VALUES (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// update account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`, t.Amount, t.AccountId)
	}

	// rollback if error
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected server error")
	}

	// commit transaction
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while committing transaction: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected server error")
	}

	// retrieve last transaction id
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while retrieving the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected server error")
	}

	// get latest account info
	account, appErr := db.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}

	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount

	return &t, nil
}

func (db AccountRepositoryDB) FindBy(accountId string) (*Account, *errs.AppError) {
	var account Account

	sqlGetAccount := `SELECT account_id, customer_id, opening_date, account_type, amount FROM accounts WHERE account_id = ?`

	err := db.Client.Get(&account, sqlGetAccount, accountId)
	if err != nil {
		logger.Error("Error while fetching account information: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected server error")
	}

	return &account, nil
}
