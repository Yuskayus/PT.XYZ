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

func (r *KonsumenRepository) GetKonsumenByID(id uint) (*domain.Konsumen, error) {
	var konsumen domain.Konsumen
	err := r.DB.First(&konsumen, id).Error
	return &konsumen, err
}

func (r *KonsumenRepository) UpdateLimit(id uint, tenor int, amount float64) error {
	konsumen, err := r.GetKonsumenByID(id)
	if err != nil {
		return err
	}

	switch tenor {
	case 1:
		konsumen.Limit1 -= amount
	case 2:
		konsumen.Limit2 -= amount
	case 3:
		konsumen.Limit3 -= amount
	case 6:
		konsumen.Limit6 -= amount
	}

	return r.DB.Save(&konsumen).Error
}
