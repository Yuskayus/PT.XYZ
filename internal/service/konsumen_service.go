// internal/service/konsumen_service.go
package service

import (
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
