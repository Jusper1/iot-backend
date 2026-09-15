package repositories

import(
	"iot-backend/internal/models"

	"gorm.io/gorm"
)

type ManualRepository struct {
	DB * gorm.DB
}

func NewManualRepository(db *gorm.DB) * ManualRepository {
	return &ManualRepository{
		DB: db,
	}
}

func (r *ManualRepository) Create(
	order *models.ManualOrder,
) error {
	return  r.DB.Create(order).Error
}

func (r *ManualRepository) FindAll() ([]models.ManualOrder,error) {
	var orders []models.ManualOrder

	err := r.DB.
	Order("id DESC").
	Find(&orders).Error

	return orders,err
}

func (r *ManualRepository) FindByID(id uint64) (*models.ManualOrder,error) {
	var order models.ManualOrder

	err := r.DB.
	Where("id = ?",id).
	First(&order).Error

	if err != nil {
		return  nil,err
	}

	return &order, nil
}

func (r *ManualRepository) Update(order *models.ManualOrder,
	) error {
		return r.DB.
		Model(&models.ManualOrder{}).
		Where("id = ?",order.ID).
		Updates(order).Error
	}

func (r *ManualRepository) Delete(id uint64)error {
	return r.DB.
	Delete(&models.ManualOrder{}, id).
	Error
}