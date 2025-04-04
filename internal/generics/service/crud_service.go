package service

import (
	"context"
	"time"

	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/i_crud"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ProductService define las operaciones de negocio para Product.

type Trackable interface {
	SetID(id primitive.ObjectID)
	SetCreationDate(t time.Time)
	SetUpdateDate(t time.Time)
}

type crudService[T any] struct {
	repo i_crud.CRUDRepository[T]
}

// NewProductService crea una nueva instancia de ProductService inyectando el repositorio.
func NewCRUDService[T mql_request_filter.EntityModel](repo i_crud.CRUDRepository[T]) i_crud.CRUDService[T] {
	return &crudService[T]{
		repo: repo,
	}
}

func (s *crudService[T]) Insert(ctx context.Context, entity *T) error {
	return s.repo.Insert(ctx, entity)
}

func (s *crudService[T]) FindByID(ctx context.Context, id string) (*T, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *crudService[T]) Update(ctx context.Context, id string, entity *T) error {
	return s.repo.Update(ctx, id, entity)
}

func (s *crudService[T]) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *crudService[T]) FindAll(ctx context.Context, filters map[string]interface{}, page, limit int) ([]T, int64, error) {
	// Validar y sanitizar el filtro
	var model T
	whitelist, err := any(model).(mql_request_filter.EntityModel).GetFilterWhitelist()
	if err != nil {
		return nil, 0, err
	}

	filter, err := mql_request_filter.ValidateAndSanitizeFilter(filters, whitelist)
	if err != nil {
		return nil, 0, err
	}

	return s.repo.FindAll(ctx, filter, page, limit)
}

// func (s *crudService[T]) GetPaginatedentityService[T]s(ctx context.Context, page int, limit int) ([]models.Product, int64, error) {
// 	return s.repo.FindPaginated(ctx, page, limit)
// }
