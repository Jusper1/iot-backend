package repositories

import (
	"iot-backend/internal/models"

	"gorm.io/gorm"
)

type PembelianAlatRepository struct {
	DB *gorm.DB
}

func NewPembelianAlatRepository(db *gorm.DB) *PembelianAlatRepository {
	return &PembelianAlatRepository{
		DB: db,
	}
}

func (r *PembelianAlatRepository) Create(
	data *models.PembelianAlat,
) error {
	return r.DB.Create(data).Error
}

func (r *PembelianAlatRepository) FindAll() (
	[]models.PembelianAlat,
	error,
) {
	var data []models.PembelianAlat

	err := r.DB.
		Order("id DESC").
		Find(&data).Error

	return data, err
}

func (r *PembelianAlatRepository) FindByID(
	id uint,
) (*models.PembelianAlat, error) {

	var data models.PembelianAlat

	err := r.DB.
		First(&data, id).
		Error

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *PembelianAlatRepository) Update(
	data *models.PembelianAlat,
) error {
	return r.DB.
		Model(&models.PembelianAlat{}).
		Where("id = ?",data.ID).
		Updates(data).Error
	}

func (r *PembelianAlatRepository) Delete(
	data *models.PembelianAlat,
) error {
	return r.DB.Delete(data).Error
}