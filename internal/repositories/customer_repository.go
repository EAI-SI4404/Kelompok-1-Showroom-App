package repositories

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Insert(customer *domain.Customer) (*domain.Customer, error)
	FindByEmail(email string) (*domain.Customer, error)
	FindRoleByName(name string) (*domain.Role, error)
	GetCustomer(id string) (*domain.Customer, error)
	GetAllCustomer() ([]domain.Customer, error)
}

type customerConnection struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) *customerConnection {
	return &customerConnection{db}
}

func (c *customerConnection) Insert(customer *domain.Customer) (*domain.Customer, error) {
	err := c.db.Create(&customer).Error
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *customerConnection) FindByEmail(email string) (*domain.Customer, error) {
	var customer domain.Customer
	err := c.db.Where("email = ?", email).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerConnection) FindRoleByName(name string) (*domain.Role, error) {
	if name == "" {
		name = "user"
	}

	var role domain.Role

	err := c.db.Where("name = ?", name).First(&role).Error

	if err != nil {
		return nil, err
	}

	return &role, nil
}

func (c *customerConnection) GetCustomer(id string) (*domain.Customer, error) {
	var customer domain.Customer
	err := c.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}
	return &customer, nil
}

func (c *customerConnection) GetAllCustomer() ([]domain.Customer, error) {
	var customers []domain.Customer
	err := c.db.Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

// Path: internal\repositories\customer_repository.go
