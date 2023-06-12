package utils

import (
	"time"

	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/config"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/domain"
	"github.com/EAI-SI4404/Kelompok-1-Showroom-App/internal/dto"
	"github.com/golang-jwt/jwt/v4"
)

func CreateUserToken(customer *domain.Customer) (string, error) {
	claims := dto.Claims{
		UserID: customer.ID,
		Email:  customer.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    customer.Name,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte("secret"))

	customerToken := domain.CustomerToken{
		CustomerID: customer.ID,
		Role:       customer.Role,
		Token:      signedToken,
	}

	var existingToken domain.CustomerToken

	err := config.DB.Where("customer_id = ?", customer.ID).First(&existingToken).Error

	if err != nil {
		err = config.DB.Create(&customerToken).Error
	} else {
		err = config.DB.Model(&existingToken).Updates(&customerToken).Error
	}

	if err != nil {
		return "", err
	}

	return signedToken, err

}
