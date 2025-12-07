package models

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Buildings []Building `gorm:"foreignKey:CustomerID" json:"buildings,omitempty"`
}

func (c *Customer) TableName() string {
	return "customer"
}

