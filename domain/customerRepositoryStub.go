package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{
			Id:          1001,
			Name:        "Ashish",
			City:        "a",
			Zipcode:     "b",
			DateOfBirth: "c",
			Status:      0,
		},
		{
			10011,
			"Ashiseeh",
			"a",
			"b",
			"c",
			0,
		},
	}

	return CustomerRepositoryStub{customers}
}
