package domain

import (
	"github.com/Danny0728/BankAPI/dto"
	"github.com/Danny0728/BankAPI/errs"
)

type Customer struct {
	Id      string `json:"id" db:"customer_id"`
	Name    string `json:"name" db:"name"`
	City    string `json:"city" db:"city"`
	ZipCode string `json:"zipCode" db:"zipcode"`
	DOB     string `json:"dob" db:"date_of_birth"`
	Status  string `json:"status" db:"status"`
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDto() dto.CustomerResponse {
	return dto.CustomerResponse{
		Id:      c.Id,
		Name:    c.Name,
		City:    c.City,
		ZipCode: c.ZipCode,
		DOB:     c.DOB,
		Status:  c.statusAsText(),
	}
}

// this is the secondary port in hexagonal architecture
type CustomerRepository interface {
	FindAll(*string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
	// AddCustomer(Customer) (Customer, error)
}
