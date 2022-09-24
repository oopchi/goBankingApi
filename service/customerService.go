package service

import (
	"github.com/oopchi/banking/domain"
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
)

type CustomerService interface {
	GetAllCustomer(*int) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomer(status *int) ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll(status)

	if err != nil {
		return nil, err
	}

	response := []dto.CustomerResponse{}

	for _, customer := range customers {
		response = append(response, customer.ToDto())
	}

	return response, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)

	if err != nil {
		return nil, err
	}

	response := customer.ToDto()

	return &response, err
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
