package services

import (
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/repositories"
)

type BalanceService interface {
	TopUpBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error)
	GetBalance(id string) (*domain.CutstomerBalance, error)
	UpdateBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error)
	FindCustomerByToken(token string) (*domain.Customer, error)
}

type balanceService struct {
	balanceRepository repositories.BalanceRepository
}

func NewBalanceService(balanceRepository repositories.BalanceRepository) *balanceService {
	return &balanceService{balanceRepository}
}

func (b *balanceService) TopUpBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewBalanceRepository(conn)

	newUserBalance, err := repo.Add(balance)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to add balance",
			StatusCode: 500,
		}
	}

	return newUserBalance, nil
}

func (b *balanceService) GetBalance(id string) (*domain.CutstomerBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewBalanceRepository(conn)

	balance, err := repo.GetBalance(id)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to get balance",
			StatusCode: 500,
		}
	}

	return balance, nil
}

func (b *balanceService) UpdateBalance(balance *domain.CutstomerBalance) (*domain.CutstomerBalance, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewBalanceRepository(conn)

	updatedBalance, err := repo.UpdateBalance(balance)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to update balance",
			StatusCode: 500,
		}
	}

	return updatedBalance, nil

}

func (b *balanceService) FindCustomerByToken(token string) (*domain.Customer, error) {
	conn, err := config.Connect()

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to connect to database",
			StatusCode: 500,
		}
	}

	repo := repositories.NewBalanceRepository(conn)

	customerID, err := repo.FindCustomerByToken(token)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to get customer",
			StatusCode: 500,
		}
	}

	customer, err := repo.FindCustomerByID(customerID)

	if err != nil {
		return nil, &ErrorMessage{
			Message:    "Failed to get customer",
			StatusCode: 500,
		}
	}

	return customer, nil

	
}

