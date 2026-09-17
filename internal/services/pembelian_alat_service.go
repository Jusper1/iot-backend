package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

type PembelianAlatService struct {
	Repository *repositories.PembelianAlatRepository
}

func NewPembelianAlatService(
	repository *repositories.PembelianAlatRepository,
) *PembelianAlatService {
	return &PembelianAlatService{
		Repository: repository,
	}
}

func (s *PembelianAlatService) Create(
	data *models.PembelianAlat,
) error {

	data.InstansiPerusahaan =
		strings.TrimSpace(data.InstansiPerusahaan)

	data.NamaPIC =
		strings.TrimSpace(data.NamaPIC)

	if data.InstansiPerusahaan == "" {
		return errors.New("instansi/perusahaan wajib diisi")
	}

	return s.Repository.Create(data)
}

func (s *PembelianAlatService) FindAll() (
	[]models.PembelianAlat,
	error,
) {
	return s.Repository.FindAll()
}

func (s *PembelianAlatService) FindByID(
	id uint,
) (*models.PembelianAlat, error) {
	return s.Repository.FindByID(id)
}

func (s *PembelianAlatService) Update(
	id uint,
	data *models.PembelianAlat,
) error {

	existing, err := s.Repository.FindByID(id)

	if err != nil {
		return err
	}

	data.ID = existing.ID

	data.InstansiPerusahaan =
		strings.TrimSpace(data.InstansiPerusahaan)

	data.NamaPIC =
		strings.TrimSpace(data.NamaPIC)

	if data.InstansiPerusahaan == "" {
		return errors.New("instansi/perusahaan wajib diisi")
	}

	return s.Repository.Update(data)
}

func (s *PembelianAlatService) Delete(
	id uint,
) error {

	data, err := s.Repository.FindByID(id)

	if err != nil {
		return err
	}

	return s.Repository.Delete(data)
}