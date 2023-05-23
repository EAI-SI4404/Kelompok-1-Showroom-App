package services

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	Register(user *domain.Customer) (*domain.Customer, error)
	FindRoleByName(name string) (*domain.Role, error)
}

type customerService struct {
	customerRepository repositories.CustomerRepository
}

func NewCustomerService(customerRepository repositories.CustomerRepository) *customerService {
	return &customerService{customerRepository}
}

func (c *customerService) Register(customer *domain.Customer) (*domain.Customer, error) {
	// TODO: Implement this method
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewCustomerRepository(conn)

	// Check if email already registered
	_, err = repo.FindByEmail(customer.Email)

	if err == nil {
		return nil, &ErrorMessage{
			Message:    "Email already registered",
			StatusCode: 400,
		}
	}

	// Find role
	_, err = repo.FindRoleByName(customer.Role)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Role not found",
			StatusCode: 404,
		}
	}

	// Hash password
	password, err := bcrypt.GenerateFromPassword([]byte(customer.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to hash password",
			StatusCode: 500,
		}
	}

	customer.Password = string(password)

	// Insert customer
	customer, err = repo.Insert(customer)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to insert customer",
			StatusCode: 500,
		}
	}

	return customer, nil
}

func (c *customerService) FindRoleByName(name string) (*domain.Role, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewCustomerRepository(conn)

	role, err := repo.FindRoleByName(name)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Role not found",
			StatusCode: 404,
		}
	}

	return role, nil
}
