package dto

import (
	"strings"

	"github.com/oopchi/banking/errs"
)

type NewAccountRequest struct {
	CustomerId  int     `json:"customer_id" xml:"customer_id"`
	AccountType string  `json:"account_type" xml:"account_type"`
	Amount      float64 `json:"amount" xml:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000.00")
	}

	fAccountType := strings.ToLower(strings.Trim(r.AccountType, " "))

	if fAccountType != "saving" && fAccountType != "checking" {
		return errs.NewValidationError("Account type should be checking or saving")
	}

	return nil
}
