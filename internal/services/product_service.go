package services

import (
	"context"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/repositories"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
)

// ProductService define las operaciones de negocio para Product.
type ProductService interface {
	CRUDService[models.Product]
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]models.Product, int64, error)
}

type productService struct {
	CRUDService[models.Product]
	repo repositories.ProductRepository
}

// NewProductService crea una nueva instancia de ProductService inyectando el repositorio.
func NewProductService(repo repositories.ProductRepository) ProductService {
	genService := NewCRUDService(repo)

	return &productService{
		CRUDService: genService,
		repo:        repo,
	}
	// return &productService{CRUDService: genService}
}

func (s *productService) FindAll(ctx context.Context, filterParams map[string]interface{}, page, limit int) ([]models.Product, int64, error) {
	filter, err := mql_request_filter.ValidateAndSanitizeFilter(filterParams, models.ProductFilterWhitelist)
	if err != nil {
		return nil, 0, err
	}

	return s.repo.FindAll(ctx, filter, page, limit)
}
