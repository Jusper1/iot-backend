package services

import (
	"errors"
	"strings"

	"iot-backend/internal/models"
	"iot-backend/internal/repositories"
)

type ManualService struct {
	Repository *repositories.ManualRepository
}

func NewManualService(
	repository *repositories.ManualRepository,
) *ManualService {
	return  &ManualService{
		Repository: repository,
	}
}

func (s *ManualService) Create(
	order *models.ManualOrder,
)error {

	order.Kode = strings.TrimSpace(order.Kode)
	order.DinkesPuskesmas = strings.TrimSpace(order.DinkesPuskesmas)

	if order.Kode == "" {
		return errors.New("Kode wajib Diisi")
	}

	if order.DinkesPuskesmas == ""{
		return  errors.New("dinkes/puskesmas wajbi diisi")
	}

	if order.HargaPPN < 0 {
		return errors.New("Harga_ppn tidak boleh negatif")
	}

	if order.JumlahUangMasuk < 0 {
		return errors.New(
			"Jumlah_uang_masuk tidak boleh negatif")
	}

	return s.Repository.Create(order)
}

func (s *ManualService) FindAll() ([]models.ManualOrder,error) {

	order, err := s.Repository.FindAll()

	if err != nil {
		return nil,errors.New(
			"gagal mengambil data Pemesanan_manual: " +err.Error(),
		)
	}

	return order, nil
}

func(s *ManualService) FindByID(id uint64)(
	*models.ManualOrder, error){

		if id == 0 {
			return nil, errors.New("id tidak valid")
		}

		order, err := s.Repository.FindByID(id)

		if err != nil {
			return nil, errors.New(
				"data pemesanan_manual tidak ditemukan")
		}

		return order, nil
	}

func (s *ManualService) Update(
	id uint64,
	order *models.ManualOrder)error {

		_, err := s.Repository.FindByID(id)

		if err != nil {
			return  err
		}

		order.Kode = strings.TrimSpace(order.Kode)
		order.DinkesPuskesmas = strings.TrimSpace(order.DinkesPuskesmas)

		if order.Kode == "" {
			return errors.New("Kode Wajib Diisi")
		}

		if order.DinkesPuskesmas == "" {
			return  errors.New("dinkes/puskesmas wajib diisi")
		}

		if order.HargaPPN < 0 {
			return  errors.New("Harga_ppn tidak boleh negatif")
		}

		if order.JumlahUangMasuk < 0 {
			return  errors.New("Jumlah_uang_masuk tidak boleh negatif")
		}

		return s.Repository.Update(order)
	}	

func (s *ManualService) Delete(id uint64) error{
	_, err := s.Repository.FindByID(id)

	if err != nil {
		return errors.New("data Pemesanan manual tidak ditemukan")
	}

	return s.Repository.Delete(id)
}

