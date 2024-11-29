package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ActivityLog struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserID"`
	InvoiceID *uuid.UUID `gorm:"type:uuid"`
	Invoice   *Invoice   `gorm:"foreignKey:InvoiceID"`
	Action    string    `gorm:"type:varchar(255);not null"`
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

// TableName specifies the table name for the ActivityLog model
func (ActivityLog) TableName() string {
	return "activity_logs"
}

// BeforeCreate hook to ensure UUID is set
func (a *ActivityLog) BeforeCreate(tx *gorm.DB) error {
	if a.ID == uuid.Nil {
		a.ID = uuid.New()
	}
	return nil
}