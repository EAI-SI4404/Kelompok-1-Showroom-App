package services

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/repositories"
)

type CustomerService interface {
	Register(user *domain.Customer) (*domain.Customer, error)
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) *customerService {
	return &customerService{customerRepository}
}

func (c *customerService) Register(customer *domain.Customer) (*domain.Customer, error) {
	// TODO: Implement this method
	return nil, nil
}
