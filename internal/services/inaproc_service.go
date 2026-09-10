package services

import(
	"errors"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

type InaprocService struct {
	Repository *repositories.InaprocRepository
}

func NewInaprocService(
	repository *repositories.InaprocRepository,
) *InaprocService {
	return &InaprocService{
		Repository: repository,
	}
}

func (s *InaprocService) Create(
	data *models.InaprocOrder,
) error {

	if data.Kode == "" {
		return errors.New("Kode Wajib Diisi")
	}

	if data.Qty < 0 {
		return errors.New("Qty tidak boleh negatif")
	}

	return s.Repository.Create(data)
}

func (s *InaprocService) FindAll() ([]models.InaprocOrder, error) {
	return s.Repository.FindAll()
}

func(s *InaprocService) FindByID(id uint,
	) (*models.InaprocOrder, error){
	return s.Repository.FindByID(id)
}

func(s *InaprocService) Update(
	id uint,
	data *models.InaprocOrder,
	) error {

		if data.Kode == "" {
			return errors.New("Kode Wajib Diisi")
		}

		if data.Qty < 0 {
			return  errors.New("Qty tidak boleh negatif")
		}

		_, err := s.Repository.FindByID(id)

		if err != nil {
			return errors.New("data inaproc tidak ditemukan")
		}

		return  s.Repository.Update(id,data)
	}

func (s *InaprocService) Delete(id uint,)error{

	_, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("data inaproc tidak ditemukan")
	}

	return  s.Repository.Delete(id)
}

