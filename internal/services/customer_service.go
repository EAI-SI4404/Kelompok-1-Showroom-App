package services

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

type CustomerService interface {
	Register(user *domain.Customer) (*domain.Customer, error)
	Login(email, password string) (*domain.Customer, error)
	FindRoleByName(name string) (*domain.Role, error)
	GetCustomer(id string) (*domain.Customer, error)
	GetAllCustomer() ([]domain.Customer, error)
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

func (c *customerService) Login(email, password string) (*domain.Customer, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewCustomerRepository(conn)

	user, err := repo.FindByEmail(email)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Email not found",
			StatusCode: 404,
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Password incorrect",
			StatusCode: 400,
		}
	}

	return user, nil
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

func (c *customerService) GetCustomer(id string) (*domain.Customer, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewCustomerRepository(conn)

	customer, err := repo.GetCustomer(id)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Customer not found",
			StatusCode: 404,
		}
	}

	return customer, nil
}

func (c *customerService) GetAllCustomer() ([]domain.Customer, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewCustomerRepository(conn)

	customers, err := repo.GetAllCustomer()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Customer not found",
			StatusCode: 404,
		}
	}

	return customers, nil
}
