package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CustomerToken struct {
	ID         string   `json:"id" gorm:"primaryKey, type:uuid, default:uuid_generate_v4()"`
	CustomerID string   `json:"customer_id"`
	Customer   Customer `json:"customer" gorm:"foreignKey:CustomerID;references:ID"`
	Token      string   `json:"token"`
}

func (b *CustomerToken) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.NewString()
	return
}
