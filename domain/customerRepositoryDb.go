package domain

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/oopchi/banking/errs"
	"github.com/oopchi/banking/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status *int) ([]Customer, *errs.AppError) {
	var err error
	customers := []Customer{}
	if status != nil {
		findAllSql := `
		select 
			customer_id,
			name,
			city,
			zipcode,
			date_of_birth,
			status
		from customers
		where status = ?
		`

		err = d.client.Select(&customers, findAllSql, *status)
	} else {

		findAllSql := `
		select 
			customer_id,
			name,
			city,
			zipcode,
			date_of_birth,
			status
		from customers
		`
		err = d.client.Select(&customers, findAllSql)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {

	customerSql := `
	select 
		customer_id,
		name,
		city,
		zipcode,
		date_of_birth,
		status
	from customers
	where customer_id = ?
	`
	var c Customer

	err := d.client.Get(&c, customerSql, id)

	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		}

		logger.Error("Error while scanning customers " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	return &c, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {

	return CustomerRepositoryDb{
		client: dbClient,
	}
}
