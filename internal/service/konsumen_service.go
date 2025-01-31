// internal/service/konsumen_service.go
package service

import (
	"errors"

	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"github.com/yuskayus/pt-xyz-multifinance/internal/repository"
)

type KonsumenService struct {
	Repo *repository.KonsumenRepository
}

// GetAll - Mendapatkan semua data konsumen
func (s *KonsumenService) GetAll() ([]domain.Konsumen, error) {
	return s.Repo.GetAll()
}

// Create - Menyimpan konsumen baru
func (s *KonsumenService) Create(konsumen domain.Konsumen) error {
	return s.Repo.Create(konsumen)
}

func (s *KonsumenService) ProcessTransaction(konsumenID uint, tenor int, amount float64) error {
	konsumen, err := s.Repo.GetKonsumenByID(konsumenID)
	if err != nil {
		return errors.New("konsumen tidak ditemukan")
	}

	// Pengecekan limit untuk setiap tenor
	switch tenor {
	case 1:
		if konsumen.Limit1 < amount {
			return errors.New("limit tidak mencukupi untuk tenor 1 bulan")
		}
		konsumen.Limit1 -= amount // Mengurangi Limit1 setelah transaksi berhasil
	case 2:
		if konsumen.Limit2 < amount {
			return errors.New("limit tidak mencukupi untuk tenor 2 bulan")
		}
		konsumen.Limit2 -= amount // Mengurangi Limit2 setelah transaksi berhasil
	case 3:
		if konsumen.Limit3 < amount {
			return errors.New("limit tidak mencukupi untuk tenor 3 bulan")
		}
		konsumen.Limit3 -= amount // Mengurangi Limit3 setelah transaksi berhasil
	case 6:
		if konsumen.Limit6 < amount {
			return errors.New("limit tidak mencukupi untuk tenor 6 bulan")
		}
		konsumen.Limit6 -= amount // Mengurangi Limit6 setelah transaksi berhasil
	default:
		return errors.New("tenor tidak valid")
	}

	// Simpan perubahan limit ke database
	return s.Repo.UpdateLimit(konsumenID, tenor, amount)
}
