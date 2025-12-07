package models

import (
	"time"
)

type OPCServer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `json:"name"`
	Host        string    `json:"host"`
	Port        int       `json:"port"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (os *OPCServer) TableName() string {
	return "opc_server"
}

type OPCTag struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ServerID    uint      `gorm:"index" json:"server_id"`
	Name        string    `json:"name"`
	Path        string    `json:"path"`
	DataTypeID uint      `gorm:"index" json:"data_type_id"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// Relations
	Server   OPCServer   `gorm:"foreignKey:ServerID" json:"server,omitempty"`
	DataType OPCDatatype `gorm:"foreignKey:DataTypeID" json:"data_type,omitempty"`
}

func (ot *OPCTag) TableName() string {
	return "opc_tag"
}

type OPCDatatype struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"uniqueIndex" json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (od *OPCDatatype) TableName() string {
	return "opc_datatype"
}

