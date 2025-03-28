package i_crud

import (
	"context"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
)

type CRUDRepository[T any] interface {
	Insert(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, filter map[string]interface{}, page, limit int) ([]T, int64, error)
}

type CRUDService[T mql_request_filter.EntityModel] interface {
	Insert(ctx context.Context, entity *T) error
	FindByID(ctx context.Context, id string) (*T, error)
	Update(ctx context.Context, id string, entity *T) error
	Delete(ctx context.Context, id string) error
	FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]T, int64, error)
}
