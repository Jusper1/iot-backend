package repositories

import (
	"iot-backend/internal/models"

	"gorm.io/gorm"
)

type SPJRepository struct {
	DB *gorm.DB
}

func NewSPJRepository(db *gorm.DB) *SPJRepository {
	return &SPJRepository{
		DB: db,
	}
}

func (r *SPJRepository) Create(data *models.SPJOrder) error {
	return r.DB.Create(data).Error
}

func (r *SPJRepository) FindAll() ([]models.SPJOrder, error) {
	var data []models.SPJOrder

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *SPJRepository) FindByID(id uint) (*models.SPJOrder, error) {
	var data models.SPJOrder

	err := r.DB.
		Where("id = ?", id).
		First(&data).Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *SPJRepository) Update(data *models.SPJOrder) error {
	return r.DB.Save(data).Error
}

func (r *SPJRepository) Delete(id uint) error {
	return r.DB.Delete(&models.SPJOrder{}, id).Error
}