package dto

type CustomerResponse struct {
	Id          int    `json:"customer_id" xml:"customer_id"`
	Name        string `json:"customer_name" xml:"customer_name"`
	City        string `json:"city" xml:"city"`
	Zipcode     string `json:"zipcode" xml:"zipcode"`
	DateOfBirth string `json:"date_of_birth" xml:"date_of_birth"`
	Status      string `json:"status" xml:"status"`
}
