package repositories

import (
	"iot-backend/internal/models"

	"gorm.io/gorm"
)

type AdminRepository struct {
	DB *gorm.DB
}

func NewAdminRepository(db *gorm.DB) *AdminRepository {
	return &AdminRepository{
		DB: db,
	}
}

func (r *AdminRepository) FindByEmail(
	email string,
) (*models.AdminUser, error) {

	var user models.AdminUser

	err := r.DB.
		Where("email = ?", email).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AdminRepository) findByID(
	id uint,
) (*models.AdminUser, error) {
	var user models.AdminUser

	err := r.DB.
		Where("id = ?", id).
		First(&user).
		Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AdminRepository) Create(
	user *models.AdminUser,
) error {

	return r.DB.Create(user).Error
}

func (r *AdminRepository) UpdatePassword(
	id uint,
	PasswordHash string,
) error {
	return r.DB.
	Model(&models.AdminUser{}).
		Where("id = ?", id).
		Update("password", PasswordHash).
		Error
}
