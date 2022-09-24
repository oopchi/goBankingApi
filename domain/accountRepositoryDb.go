package domain

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	"github.com/oopchi/banking/errs"
	"github.com/oopchi/banking/logger"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (d AccountRepositoryDb) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := `
		insert into accounts (
			customer_id,
			opening_date,
			account_type,
			amount,
			status
		) values (
			?,
			?,
			?,
			?,
			?
		)
	`

	result, err := d.client.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)

	if err != nil {
		logger.Error("Error while creating new account: " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting last insert id for new account: " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	a.AccountId = int(id)

	return &a, nil
}

func (d AccountRepositoryDb) FindBy(id int) (*Account, *errs.AppError) {
	sq := `
	select
		account_id,
		customer_id,
		opening_date,
		account_type,
		amount,
		status
	from accounts
	where account_id = ?
	`

	var a Account

	if err := d.client.Get(&a, sq, id); err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning accounts: " + err.Error())
			return nil, errs.NewNotFoundError("Account not found")
		}

		logger.Error("Error while scanning accounts: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &a, nil
}

func (d AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := d.client.Begin()

	if err != nil {
		logger.Error("Error while starting a new transaction for bank account transaction: " + err.Error())

		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	query := `
		insert into transactions(
			account_id,
			amount,
			transaction_type,
			transaction_date
		) values(
			?,
			?,
			?,
			?
		)
	`
	result, err := tx.Exec(
		query,
		t.AccountId,
		t.Amount,
		t.TransactionType,
		t.TransactionDate,
	)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error("Error while rolling back failed transaction: " + err.Error())

			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		logger.Error("Error while inserting new transaction into transactions table: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	if t.IsWithdrawal() {
		query = `
			update accounts
			set amount = amount - ?
			where account_id = ?
		`
	} else {
		query = `
			update accounts
			set amount = amount + ?
			where account_id = ?
		`
	}

	_, err = tx.Exec(
		query,
		t.Amount,
		t.AccountId,
	)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error("Error while rolling back failed transaction: " + err.Error())

			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		logger.Error("Error while updating account's balance: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	err = tx.Commit()

	if err != nil {
		if err := tx.Rollback(); err != nil {
			logger.Error("Error while rolling back failed transaction: " + err.Error())

			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

		logger.Error("Error while committing transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	transactionId, err := result.LastInsertId()

	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	account, appErr := d.FindBy(t.AccountId)

	if appErr != nil {
		return nil, appErr
	}

	t.Id = int(transactionId)
	t.Amount = account.Amount

	return &t, nil
}

func NewAccountRepositoryDb(dbClient *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{
		client: dbClient,
	}
}
