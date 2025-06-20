package models

import (
	"time"

	"gorm.io/gorm"
)

// Transaction represents a financial transaction in the double-entry ledger system
type Transaction struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Description string         `gorm:"size:255" json:"description"`
	Amount      float64        `gorm:"not null" json:"amount"`
	Date        time.Time      `gorm:"not null" json:"date"`
	Status      string         `gorm:"size:50;default:'pending'" json:"status"` // pending, completed, failed, reversed
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Entries     []Entry        `gorm:"foreignKey:TransactionID" json:"entries"`
}
