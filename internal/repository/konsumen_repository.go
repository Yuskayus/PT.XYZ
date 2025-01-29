package repository

import (
	"fmt"

	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"gorm.io/gorm"
)

type KonsumenRepository struct {
	DB *gorm.DB
}

func (r *KonsumenRepository) GetAll() ([]domain.Konsumen, error) {
	var konsumens []domain.Konsumen
	err := r.DB.Find(&konsumens).Error
	return konsumens, err
}

func (r *KonsumenRepository) Create(konsumen domain.Konsumen) error {
	// Cek apakah NIK sudah ada
	var existingKonsumen domain.Konsumen
	if err := r.DB.Where("nik = ?", konsumen.NIK).First(&existingKonsumen).Error; err == nil {
		return fmt.Errorf("NIK already exists")
	}

	// Jika NIK tidak ada, lanjutkan untuk membuat konsumen
	return r.DB.Create(&konsumen).Error
}
