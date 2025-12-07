package models

import (
	"time"

	"gorm.io/gorm"
)

type ProjectSection struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Buildings []Building `gorm:"many2many:project_section_building;" json:"buildings,omitempty"`
}

func (ps *ProjectSection) TableName() string {
	return "project_section"
}

