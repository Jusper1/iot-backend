package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

type SPJService struct {
	Repository *repositories.SPJRepository
}

func NewSPJService(repository *repositories.SPJRepository) *SPJService {
	return &SPJService{
		Repository: repository,
	}
}

func (s *SPJService) Create(data *models.SPJOrder) error {
	data.NamaDinkesPKM = strings.TrimSpace(data.NamaDinkesPKM)

	if data.NamaDinkesPKM == "" {
		return errors.New("nama dinkes/pkm wajib diisi")
	}

	return s.Repository.Create(data)
}

func (s *SPJService) FindAll() ([]models.SPJOrder, error) {
	return s.Repository.FindAll()
}

func (s *SPJService) FindByID(id uint) (*models.SPJOrder, error) {
	return s.Repository.FindByID(id)
}

func (s *SPJService) Update(id uint, data *models.SPJOrder) error {
	existing, err := s.Repository.FindByID(id)
	if err != nil {
		return err
	}

	data.ID = existing.ID

	data.NamaDinkesPKM = strings.TrimSpace(data.NamaDinkesPKM)

	if data.NamaDinkesPKM == "" {
		return errors.New("nama dinkes/pkm wajib diisi")
	}

	return s.Repository.Update(data)
}

func (s *SPJService) Delete(id uint) error {
	_, err := s.Repository.FindByID(id)

	if err != nil {
		return err
	}

	return s.Repository.Delete(id)
}