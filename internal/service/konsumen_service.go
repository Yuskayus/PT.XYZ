package service

import (
	"github.com/yuskayus/pt-xyz-multifinance/internal/domain"
	"github.com/yuskayus/pt-xyz-multifinance/internal/repository"
)

type KonsumenService struct {
	Repo *repository.KonsumenRepository
}

func (s *KonsumenService) GetAll() ([]domain.Konsumen, error) {
	return s.Repo.GetAll()
}

func (s *KonsumenService) Create(input domain.Konsumen) error {
	return s.Repo.Create(input)
}
