package repositories

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"gorm.io/gorm"
)

type BalanceRepository interface {
	Add(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error)
	GetBalance(id string) (*domain.CutstomerBalance, error)
	UpdateBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error)
	FindCustomerByID(id string) (*domain.Customer, error)
	FindCustomerByToken(token string) (string, error)
}

type balanceRepository struct {
	db *gorm.DB
}

func NewBalanceRepository(db *gorm.DB) *balanceRepository {
	return &balanceRepository{db}
}

func (c *balanceRepository) Add(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error) {
	err := c.db.Create(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *balanceRepository) GetBalance(id string) (*domain.CutstomerBalance, error) {
	var balance domain.CutstomerBalance
	err := c.db.Where("customer_id = ?", id).First(&balance).Error
	if err != nil {
		return nil, err
	}
	return &balance, nil
}

func (c *balanceRepository) UpdateBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error) {
	err := c.db.Save(&balance).Error
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (c *balanceRepository) FindCustomerByID(id string) (*domain.Customer, error) {
	var customer domain.Customer
	err := c.db.Where("id = ?", id).First(&customer).Error
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (c *balanceRepository) FindCustomerByToken(token string) (string, error) {
	var customer domain.CustomerToken

	err := c.db.Where("token = ?", token).First(&customer).Error

	if err != nil {
		return "", err
	}

	return customer.CustomerID, nil
}
