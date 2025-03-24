package product

import (
	"context"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
)

// ProductService define las operaciones de negocio para Product.
type ProductService interface {
	service.CRUDService[Product]
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]Product, int64, error)
}

type productService struct {
	service.CRUDService[Product]
	repo ProductRepository
}

// NewProductService crea una nueva instancia de ProductService inyectando el repositorio.
func NewProductService(repo ProductRepository) ProductService {
	genService := service.NewCRUDService(repo)

	return &productService{
		CRUDService: genService,
		repo:        repo,
	}
	// return &productService{CRUDService: genService}
}

func (s *productService) FindAll(ctx context.Context, filterParams map[string]interface{}, page, limit int) ([]Product, int64, error) {
	filter, err := mql_request_filter.ValidateAndSanitizeFilter(filterParams, ProductFilterWhitelist)
	if err != nil {
		return nil, 0, err
	}

	return s.repo.FindAll(ctx, filter, page, limit)
}
