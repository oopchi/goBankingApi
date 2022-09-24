package domain

import (
	"github.com/oopchi/banking/dto"
	"github.com/oopchi/banking/errs"
	"github.com/oopchi/banking/util"
)

type Customer struct {
	Id          int `db:"customer_id"`
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      int
}

type CustomerRepository interface {
	FindAll(*int) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      util.StatusText(c.Status),
	}
}
