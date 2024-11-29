package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaymentDetails struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	InvoiceID       uuid.UUID `gorm:"type:uuid;not null"`
	Invoice         Invoice   `gorm:"foreignKey:InvoiceID"`
	AccountName     string    `gorm:"type:varchar(100);not null"`
	AccountNumber   string    `gorm:"type:varchar(50);not null"`
	BankName        string    `gorm:"type:varchar(100)"`
	BankAddress     string    `gorm:"type:varchar(255)"`
	RoutingNumber   string    `gorm:"type:varchar(50)"`
	PaymentDueDate  time.Time `gorm:"not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

func (PaymentDetails) TableName() string {
	return "payment_details"
}

func (p *PaymentDetails) BeforeCreate(tx *gorm.DB) error {
	if p.ID == uuid.Nil {
		p.ID = uuid.New()
	}
	return nil
}

