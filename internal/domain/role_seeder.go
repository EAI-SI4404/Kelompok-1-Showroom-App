package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedRoles(db *gorm.DB) error {
	// Create admin role and user role
	roles := []Role{
		{
			ID:   uuid.NewString(),
			Name: "admin",
		},
		{
			ID:   uuid.NewString(),
			Name: "customer",
		},
	}

	for _, role := range roles {
		// Check if role already exists
		existingRole := Role{}
		result := db.Where("name = ?", role.Name).First(&existingRole)
		if result.Error == nil {
			// Role with the given name already exists, skip seeding
			continue
		}

		if err := db.Create(&role).Error; err != nil {
			return err
		}
	}

	return nil
}
