package domain

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeedCustomers(db *gorm.DB) error {
	// Check if customer already exists
	existingCustomer := Customer{}
	result := db.Where("name = ?", "jhon").First(&existingCustomer)
	if result.Error == nil {
		// Customer with the given name already exists, skip seeding
		return nil
	}

	id := uuid.NewString()
	password := "password"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	customer := Customer{
		ID:       id,
		Name:     "jhon",
		Email:    "jhon@gmail.com",
		Phone:    "081234567890",
		Password: string(hashedPassword),
		Address:  "Street 1",
		City:     "San Francisco",
		Province: "California",
		ZipCode:  "12345",
		Role:     "customer",
	}

	if err := db.Create(&customer).Error; err != nil {
		return err
	}

	return nil
}
