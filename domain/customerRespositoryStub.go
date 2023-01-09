package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

// the adapter which implements the secondary port (repository)
func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customers, nil
}

// func (c CustomerRepositoryStub) Add() ([]Customer, error) {
// 	dhoni := Customer{"007", "dhoni", "chennai", "421301", "7 July", "Active"}
// 	c.customers = append(c.customers, dhoni)
// 	return c.customers, nil
// }

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"007", "Yash", "Mumbai", "421301", "22 April", "Active"},
		{"008", "Mahi", "Kalyan", "421301", "12 May", "InActive"},
	}

	return CustomerRepositoryStub{customers: customers}
}
