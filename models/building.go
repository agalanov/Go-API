package models

import (
	"time"

	"gorm.io/gorm"
)

type Building struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ExtID       string    `json:"ext_id"`
	CustomerID  uint      `gorm:"index" json:"customer_id"`
	Address     string    `json:"address"`
	Description string    `json:"description"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Created     time.Time `json:"created"`
	Modified    time.Time `json:"modified"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Customer      Customer           `gorm:"foreignKey:CustomerID" json:"customer,omitempty"`
	Equipments    []Equipment        `gorm:"foreignKey:BuildingID" json:"equipments,omitempty"`
	Floors        []Floor            `gorm:"foreignKey:BuildingID" json:"floors,omitempty"`
	ProjectSections []ProjectSection `gorm:"many2many:project_section_building;" json:"project_sections,omitempty"`
}

func (b *Building) TableName() string {
	return "building"
}

func (b *Building) NotDeleted() *gorm.DB {
	return DB.Where("deleted_at IS NULL")
}

