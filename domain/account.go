package domain

import (
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
)

type Account struct {
	AccountId   int    `db:"account_id"`
	CustomerId  int    `db:"customer_id"`
	OpeningDate string `db:"opening_date"`
	AccountType string `db:"account_type"`
	Amount      float64
	Status      int
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	FindBy(int) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (a Account) ToNewAcountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

func (a Account) CanWithdraw(amount float64) bool {
	return a.Amount > amount
}
