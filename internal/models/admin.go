package models

import "time"

type AdminUser struct {
	ID uint `json:"id" gorm:"primaryKey"`

	Nama string `json:"nama"`

	Email string `json:"email" gorm:"uniqueIndex"`

	PasswordHash string `json:"-"`

	Role string `json:"role"`

	Flag int `json:"flag"`

	CreatedAt time.Time `json:"created_at"`

	UpdatedAt time.Time `json:"updated_at"`
}

func (AdminUser) TableName() string {
	return "admin_users"
}