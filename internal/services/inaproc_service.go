package services

import (
	"errors"
	"strings"

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

	data.Kode = strings.TrimSpace(data.Kode)

	if data.Kode == "" {
		return errors.New("Kode Wajib Diisi")
	}

	if data.Qty < 0 {
		return errors.New("Qty tidak boleh negatif")
	}

	if data.HargaPPN < 0 {
		return errors.New("harga_ppn tidak boleh kurang dari 0")
	}

	if data.JumlahUangMasuk < 0 {
		return errors.New(
			"jumlah_uang_masuk tidak boleh kurang dari 0",
		)
	}
	
	if err := s.Repository.Create(data); err != nil {
		return errors.New(
			"gagal menyimpan data inaproc: " + err.Error(),
		)
	}

	return s.Repository.Create(data)
}

func (s *InaprocService) FindAll() ([]models.InaprocOrder, error) {


	data, err := s.Repository.FindAll()

	if err != nil {
		return nil, errors.New(
			"gagal mengambil data inaproc: " + err.Error(),
		)
	}

	return data, nil
}

func(s *InaprocService) FindByID(id uint,
	) (*models.InaprocOrder, error){
	
	if id == 0 {
			return nil, errors.New("id tidak valid")
		}

		data, err := s.Repository.FindByID(id)

		if err != nil {
			return nil, errors.New("data inaproc tidak ditemukan")
		}

	return data, nil

}

func(s *InaprocService) Update(
	id uint,
	data *models.InaprocOrder,
	) error {

		if id == 0 {
		return errors.New("id tidak valid")
		}

		data.Kode = strings.TrimSpace(data.Kode)

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

		if data.HargaPPN < 0 {
		return errors.New("harga_ppn tidak boleh kurang dari 0")
		}

		if data.JumlahUangMasuk < 0 {
		return errors.New(
			"jumlah_uang_masuk tidak boleh kurang dari 0",
		)
		}

		if err := s.Repository.Update(id, data); err != nil {
		return errors.New(
			"gagal mengupdate data inaproc: " + err.Error(),
		)
		}

		return  nil
	}

func (s *InaprocService) Delete(id uint,)error{

	if id == 0 {
		return errors.New("id tidak valid")
	}

	_, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("data inaproc tidak ditemukan")
	}

	if err := s.Repository.Delete(id); err != nil {
		return errors.New(
			"gagal menghapus data inaproc: " + err.Error(),
		)
	}

	return  s.Repository.Delete(id)
}

