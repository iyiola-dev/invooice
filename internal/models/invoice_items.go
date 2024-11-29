package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type InvoiceItem struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	InvoiceID   uuid.UUID `gorm:"type:uuid;not null"`
	Invoice     Invoice   `gorm:"foreignKey:InvoiceID"`
	Description string    `gorm:"type:text;not null"`
	Quantity    int       `gorm:"not null"`
	UnitPrice   float64   `gorm:"type:decimal(10,2);not null"`
	Amount      float64   `gorm:"type:decimal(10,2);not null"`
}

func (InvoiceItem) TableName() string {
	return "invoice_items"
}

func (i *InvoiceItem) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}