package models

import (
	"time"
)

type TokenType int

const (
	TokenTypeQueryParam TokenType = iota + 1
	TokenTypeOAuthAccess
	TokenTypeOAuthRefresh
	TokenTypeJWT
)

type Token struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	TypeID    int       `gorm:"not null" json:"type_id"`
	Value     string    `gorm:"uniqueIndex;not null" json:"value"`
	ExpiresAt *time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Relations
	User User `gorm:"foreignKey:UserID" json:"-"`
}

func (t *Token) TableName() string {
	return "token"
}

func (t *Token) IsExpired() bool {
	if t.ExpiresAt == nil {
		return false
	}
	return time.Now().After(*t.ExpiresAt)
}

func (t *Token) IsLive() bool {
	return !t.IsExpired()
}

