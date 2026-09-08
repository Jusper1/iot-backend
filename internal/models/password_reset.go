package models

import "time"

type PasswordResetToken struct {
	ID uint `json:"id" gorm:"primaryKey"`

	UserID uint `json:"user_id"`

	TokenHash string `json:"-"`

	ExpiresAt time.Time `json:"expires_at"`

	UsedAt *time.Time `json:"used_at"`

	CreatedAt time.Time `json:"created_at"`
}

func (PasswordResetToken) TableName() string {
	return "password_reset_tokens"
}