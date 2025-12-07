package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	Username     string    `gorm:"uniqueIndex;not null" json:"username"`
	PasswordHash string    `gorm:"column:password_hash;not null" json:"-"`
	Name         string    `json:"name"`
	Email        string    `gorm:"uniqueIndex;not null" json:"email"`
	Phone        string    `json:"phone"`
	Comment      string    `json:"comment"`
	Locale       string    `gorm:"default:'en_US'" json:"locale"`
	StatusID     int       `gorm:"default:1" json:"status_id"`
	TypeID       int       `gorm:"default:1" json:"type_id"`
	IsAdmin      bool      `gorm:"default:false" json:"is_admin"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relations
	Tokens []Token `gorm:"foreignKey:UserID" json:"-"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}

func (u *User) IsSuperuser() bool {
	return u.IsAdmin
}

