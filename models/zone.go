package models

import (
	"time"

	"gorm.io/gorm"
)

type Zone struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	FloorID     *uint     `gorm:"index" json:"floor_id"`
	BuildingID  uint      `gorm:"index" json:"building_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Floor      *Floor      `gorm:"foreignKey:FloorID" json:"floor,omitempty"`
	Building   Building    `gorm:"foreignKey:BuildingID" json:"building,omitempty"`
	Equipments []Equipment `gorm:"foreignKey:ZoneID" json:"equipments,omitempty"`
}

func (z *Zone) TableName() string {
	return "zone"
}

type ZoneGroup struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (zg *ZoneGroup) TableName() string {
	return "zone_groups"
}

