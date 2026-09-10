package repositories

import (
	"gorm.io/gorm"

	"iot-backend/internal/models"
)

type InaprocRepository struct{
	 DB *gorm.DB
}

func NewInaprocRepository(db *gorm.DB) * InaprocRepository {
	return &InaprocRepository{
		DB: db,
	}
}

func (r *InaprocRepository) Create(
	data *models.InaprocOrder,
) error {
	return r.DB.Create(data).Error
}

func (r *InaprocRepository) FindAll() ([]models.InaprocOrder, error) {
	var data []models.InaprocOrder

	err := r.DB.
	Order("id DESC").
	Find(&data).Error

	return  data, err
}

func (r *InaprocRepository) FindByID(id uint,) (*models.InaprocOrder, error){
	var data models.InaprocOrder

	err := r.DB.
	First(&data, id).Error

	if err != nil{
		return nil,err
	}

	return &data, nil
}

func (r *InaprocRepository) Update(id uint, data *models.InaprocOrder,
	) error {

		return r.DB.
		Model(&models.InaprocOrder{}).
		Where("id = ?",id).
		Updates(data).Error
}

func (r *InaprocRepository) Delete(id uint,)error {
	return r.DB.
		Delete(&models.InaprocOrder{}, id).
		Error
}