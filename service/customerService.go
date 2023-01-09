package service

import (
	"github.com/Danny0728/BankAPI/domain"
	"github.com/Danny0728/BankAPI/dto"
	"github.com/Danny0728/BankAPI/errs"
)

// this is the primary port(implmented by business logic (domain)) in hexagonal architecture
type CustomerService interface {
	GetAllCustomer(*string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
	// AddCustomer(domain.Customer) (domain.Customer, error)
}

// dependency of the secondary port
type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

// this is the business logic which implements the primary port and has the dependency of secondary port
func (r DefaultCustomerService) GetAllCustomer(status *string) ([]dto.CustomerResponse, *errs.AppError) {
	c, err := r.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	var response []dto.CustomerResponse

	for i := 0; i < len(c); i++ {
		tempCustomer := c[i].ToDto()
		response = append(response, tempCustomer)
	}
	return response, nil //this connects primary port to secondary port
}

func (r DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := r.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := c.ToDto()
	return &response, nil //this connects primary port to secondary port
}

// func (r DefaultCustomerService) AddCustomer(domain.Customer) (domain.Customer, error) {
// 	return r.repo.AddCustomer(domain.Customer{})
// }

// to initiate the defaultcustomerservice this is the helper function
func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo: repository}
}
