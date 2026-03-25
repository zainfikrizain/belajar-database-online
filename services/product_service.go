package services

import (
	"Tugas-2/models"
	"Tugas-2/repositories"
	"context"
	"fmt"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]models.Product, error) {
	return s.repo.GetAll(ctx)
}

func (s *ProductService) Create(ctx context.Context, p *models.Product) error {
	if err := s.repo.Create(ctx, p); err != nil {
		return fmt.Errorf("create failed: %w", err) // ✅ keep original error
	}
	return nil
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}
