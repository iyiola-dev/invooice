package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Invoice struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID        uuid.UUID `gorm:"type:uuid;not null"`
	User          User      `gorm:"foreignKey:UserID"`
	CustomerID    uuid.UUID `gorm:"type:uuid;not null"`
	Customer      Customer  `gorm:"foreignKey:CustomerID"`
	InvoiceNumber string    `gorm:"type:varchar(50);not null"`
	IssueDate     time.Time `gorm:"not null"`
	DueDate       time.Time `gorm:"not null"`
	Currency      string    `gorm:"type:varchar(10);not null"`
	SubTotal      float64   `gorm:"type:decimal(10,2);not null"`
	Discount      float64   `gorm:"type:decimal(10,2)"`
	TotalAmount   float64   `gorm:"type:decimal(10,2);not null"`
	Status        string    `gorm:"type:varchar(20);default:'pending'"`
	Note          string    `gorm:"type:text"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`
	Items         []InvoiceItem `gorm:"foreignKey:InvoiceID"`
}

func (Invoice) TableName() string {
	return "invoices"
}

func (i *Invoice) BeforeCreate(tx *gorm.DB) error {
	if i.ID == uuid.Nil {
		i.ID = uuid.New()
	}
	return nil
}