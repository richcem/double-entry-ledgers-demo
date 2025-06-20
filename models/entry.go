package models

import (
	"time"

	"gorm.io/gorm"
)

// Entry represents a single accounting entry (debit or credit) in a transaction
type Entry struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	TransactionID uint           `gorm:"not null" json:"transaction_id"`
	AccountID     uint           `gorm:"not null" json:"account_id"`
	Amount        float64        `gorm:"not null" json:"amount"`       // Positive for debit, negative for credit
	Type          string         `gorm:"size:10;not null" json:"type"` // debit or credit
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relationships
	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	Account     Account     `gorm:"foreignKey:AccountID" json:"account"`
}
