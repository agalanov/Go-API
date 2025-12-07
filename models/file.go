package models

import (
	"time"

	"gorm.io/gorm"
)

type EntityType int

const (
	EntityTypeBuilding EntityType = iota + 1
	EntityTypeFloor
	EntityTypeEquipment
)

type File struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	EntityID   uint      `gorm:"index" json:"entity_id"`
	EntityType int       `gorm:"index" json:"entity_type"`
	Name       string    `json:"name"`
	Path       string    `json:"path"`
	Size       int64     `json:"size"`
	MimeType   string    `json:"mime_type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

func (f *File) TableName() string {
	return "file"
}

