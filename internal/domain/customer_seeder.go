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

	// Normal User
	normalUser := Customer{
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

	if err := db.Create(&normalUser).Error; err != nil {
		return err
	}

	// Admin User
	id = uuid.NewString()
	password = "adminpassword"

	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	adminUser := Customer{
		ID:       id,
		Name:     "admin",
		Email:    "admin@gmail.com",
		Phone:    "0987654321",
		Password: string(hashedPassword),
		Address:  "Street 2",
		City:     "San Francisco",
		Province: "California",
		ZipCode:  "54321",
		Role:     "admin",
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	return nil
}
