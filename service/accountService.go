package service

import (
	"os"
	"strings"
	"time"

	"github.com/oopchi/banking/domain"
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
)

type AccountService interface {
	NewAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repo domain.AccountRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError) {
	err := req.Validate()

	if err != nil {
		return nil, err
	}

	a := domain.Account{
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(os.Getenv("TIME_FORMAT")),
		AccountType: strings.ToLower(strings.Trim(req.AccountType, " ")),
		Amount:      req.Amount,
		Status:      1,
	}

	newAccount, err := s.repo.Save(a)

	if err != nil {
		return nil, err
	}

	response := newAccount.ToNewAcountResponseDto()

	return &response, nil
}

func (s DefaultAccountService) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	if req.IsTransactionTypeWithdrawal() {
		account, err := s.repo.FindBy(req.AccountId)

		if err != nil {
			return nil, err
		}

		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")
		}
	}

	t := domain.Transaction{
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: strings.ToLower(strings.Trim(req.TransactionType, " ")),
		TransactionDate: time.Now().Format(os.Getenv("TIME_FORMAT")),
	}

	transaction, err := s.repo.SaveTransaction(t)

	if err != nil {
		return nil, err
	}

	response := transaction.ToDto()

	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{
		repo: repo,
	}
}
