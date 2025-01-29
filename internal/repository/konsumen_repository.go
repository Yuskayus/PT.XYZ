// internal/repository/konsumen_repository.go
package repository

import (
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"gorm.io/gorm"
)

type KonsumenRepository struct {
	DB *gorm.DB
}

// GetAll - Mengambil semua data konsumen
func (r *KonsumenRepository) GetAll() ([]domain.Konsumen, error) {
	var konsumens []domain.Konsumen
	err := r.DB.Find(&konsumens).Error
	return konsumens, err
}

// Create - Menyimpan data konsumen baru
func (r *KonsumenRepository) Create(konsumen domain.Konsumen) error {
	return r.DB.Create(&konsumen).Error
}
