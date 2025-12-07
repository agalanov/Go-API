package models

import (
	"time"

	"gorm.io/gorm"
)

type Equipment struct {
	ID               uint      `gorm:"primaryKey" json:"id"`
	ExtID            string    `json:"ext_id"`
	EntityID         uint      `gorm:"index" json:"entity_id"`
	ZoneID           *uint     `gorm:"index" json:"zone_id"`
	BuildingID       uint      `gorm:"index" json:"building_id"`
	FloorID          *uint     `gorm:"index" json:"floor_id"`
	Description      string    `json:"description"`
	PositionX        float64   `json:"position_x"`
	PositionY        float64   `json:"position_y"`
	PositionZ        float64   `json:"position_z"`
	Enabled          bool      `gorm:"default:true" json:"enabled"`
	ProductionDate   *time.Time `json:"production_date"`
	CommissioningDate *time.Time `json:"commissioning_date"`
	Cost             float64   `json:"cost"`
	Created          time.Time `json:"created"`
	Modified         time.Time `json:"modified"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Entity   EquipmentEntity   `gorm:"foreignKey:EntityID" json:"entity,omitempty"`
	Zone     *Zone             `gorm:"foreignKey:ZoneID" json:"zone,omitempty"`
	Building Building          `gorm:"foreignKey:BuildingID" json:"building,omitempty"`
	Floor    *Floor            `gorm:"foreignKey:FloorID" json:"floor,omitempty"`
	Binds    []EquipmentBind   `gorm:"foreignKey:EquipmentID" json:"binds,omitempty"`
	Devices  []EquipmentDevice `gorm:"foreignKey:EquipmentID" json:"devices,omitempty"`
	Limits   []EquipmentLimit  `gorm:"foreignKey:EquipmentID" json:"limits,omitempty"`
}

func (e *Equipment) TableName() string {
	return "equipment"
}

type EquipmentType struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (et *EquipmentType) TableName() string {
	return "equipment_type"
}

type EquipmentEntity struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	TypeID      uint      `gorm:"index" json:"type_id"`
	SectionID   *uint     `gorm:"index" json:"section_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Type    EquipmentType   `gorm:"foreignKey:TypeID" json:"type,omitempty"`
	Section *ProjectSection `gorm:"foreignKey:SectionID" json:"section,omitempty"`
}

func (ee *EquipmentEntity) TableName() string {
	return "equipment_entity"
}

type EquipmentBind struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EquipmentID uint      `gorm:"index" json:"equipment_id"`
	BindID      uint      `gorm:"index" json:"bind_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (eb *EquipmentBind) TableName() string {
	return "equipment_bind"
}

type EquipmentDirection struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	BuildingID  uint      `gorm:"index" json:"building_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ed *EquipmentDirection) TableName() string {
	return "equipment_direction"
}

type EquipmentLimit struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EquipmentID uint      `gorm:"index" json:"equipment_id"`
	MinValue    float64   `json:"min_value"`
	MaxValue    float64   `json:"max_value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (el *EquipmentLimit) TableName() string {
	return "equipment_limit"
}

type EquipmentAttribute struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EquipmentID uint      `gorm:"index" json:"equipment_id"`
	Name        string    `json:"name"`
	Value       string    `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ea *EquipmentAttribute) TableName() string {
	return "equipment_attribute"
}

type EquipmentDevice struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EquipmentID uint      `gorm:"index" json:"equipment_id"`
	HostID      uint      `gorm:"index" json:"host_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Equipment Equipment `gorm:"foreignKey:EquipmentID" json:"equipment,omitempty"`
}

func (ed *EquipmentDevice) TableName() string {
	return "equipment_device"
}

type EquipmentState struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	EquipmentID uint      `gorm:"index" json:"equipment_id"`
	State       string    `json:"state"`
	Value       float64   `json:"value"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (es *EquipmentState) TableName() string {
	return "equipment_state"
}

