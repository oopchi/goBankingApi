package dto

import (
	"strings"

	"github.com/oopchi/banking/errs"
)

type TransactionRequest struct {
	CustomerId      int
	AccountId       int
	Amount          float64 `json:"amount" xml:"amount"`
	TransactionType string  `json:"transaction_type" xml:"transaction_type"`
}

func (r TransactionRequest) Validate() *errs.AppError {
	if r.Amount < .0 {
		return errs.NewValidationError("Amount cannot be negative.")
	}

	if !(r.IsTransactionTypeDeposit() || r.IsTransactionTypeWithdrawal()) {
		return errs.NewValidationError("Transaction can only be either withdrawal or deposit")
	}

	return nil
}

func (r TransactionRequest) IsTransactionTypeWithdrawal() bool {
	fTransactionType := strings.ToLower(strings.Trim(r.TransactionType, " "))

	return fTransactionType == "withdrawal"
}

func (r TransactionRequest) IsTransactionTypeDeposit() bool {
	fTransactionType := strings.ToLower(strings.Trim(r.TransactionType, " "))

	return fTransactionType == "deposit"
}
