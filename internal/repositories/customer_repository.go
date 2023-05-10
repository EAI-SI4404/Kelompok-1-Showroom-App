package repositories

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Insert(customer *domain.Customer) (*domain.Customer, error)
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

// Path: internal\repositories\customer_repository.go

