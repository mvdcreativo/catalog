package order

import (
	"context"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
)

type OrderService interface {
	service.CRUDService[Order]
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]Order, int64, error)
}

type orderService struct {
	service.CRUDService[Order]
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderService {
	genService := service.NewCRUDService(repo)

	return &orderService{
		CRUDService: genService,
		repo:        repo,
	}
}

func (s *orderService) FindAll(ctx context.Context, filterParams map[string]interface{}, page, limit int) ([]Order, int64, error) {
	filter, err := mql_request_filter.ValidateAndSanitizeFilter(filterParams, OrderFilterWhitelist)
	if err != nil {
		return nil, 0, err
	}

	return s.repo.FindAll(ctx, filter, page, limit)
}
