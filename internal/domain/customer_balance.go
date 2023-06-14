package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CutstomerBalance struct {
	ID         string    `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	CustomerID string    `json:"customer_id"`
	Customer   Customer  `json:"customer" gorm:"foreignKey:CustomerID;references:ID"`
	Balance    float64   `json:"balance" gorm:"default:0"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (b *CutstomerBalance) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}
