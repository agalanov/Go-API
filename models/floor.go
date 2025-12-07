package models

import (
	"time"

	"gorm.io/gorm"
)

type Floor struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BuildingID  uint      `gorm:"index" json:"building_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Level       int       `json:"level"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Building  Building    `gorm:"foreignKey:BuildingID" json:"building,omitempty"`
	Zones     []Zone      `gorm:"foreignKey:FloorID" json:"zones,omitempty"`
	Equipments []Equipment `gorm:"foreignKey:FloorID" json:"equipments,omitempty"`
}

func (f *Floor) TableName() string {
	return "floor"
}

type FloorMode struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FloorID     uint      `gorm:"index" json:"floor_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

func (fm *FloorMode) TableName() string {
	return "floor_mode"
}

