package domain

import (
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
)

type Transaction struct {
	Id              int `db:"transaction_id"`
	AccountId       int
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	Save(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		Id:              t.Id,
		AccountId:       t.AccountId,
		Amount:          t.Amount,
		TransactionType: t.TransactionType,
		TransactionDate: t.TransactionDate,
	}
}

func (t Transaction) IsWithdrawal() bool {
	return t.TransactionType == "withdrawal"
}
