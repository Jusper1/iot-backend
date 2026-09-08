package repositories

import (
	"iot-backend/internal/models"
	"time"

	"gorm.io/gorm"
)

type PasswordResetRepository  struct {
	DB *gorm.DB
}

func NewPasswordResetRepository (db *gorm.DB) *PasswordResetRepository  {
	return &PasswordResetRepository {
		DB: db,
	}
}

func (r *PasswordResetRepository ) Create(
	user *models.PasswordResetToken,
) error {
	return r.DB.Create(user).Error
}

func (r *PasswordResetRepository ) FindValidToken(
	tokenHash string,
) (*models.PasswordResetToken, error) {
	var token models.PasswordResetToken

	err := r.DB.
		Where(
			"token_hash = ? AND used_at > ? AND expires_at > ?",
			tokenHash,
			time.Now(),
		).
		First(&token).
		Error

	if err != nil {
		return nil, err
	}

	return &token, nil
}

func (r *PasswordResetRepository ) MarkUsed(
	id uint,
) error {
	now := time.Now()

	return r.DB.
		Model(&models.PasswordResetToken{}).
		Where("id = ?",id).
		Update("used_at", &now).
		Error
}
