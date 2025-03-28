package category

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/i_crud"
)

// CategoryService define las operaciones de negocio para Category.

type CategoryService interface {
	i_crud.CRUDService[Category]
}

type categoryService struct {
	i_crud.CRUDService[Category]
}

// NewCategoryService crea una nueva instancia de CategoryService inyectando el repositorio.
func NewCategoryService(repo CategoryRepository) CategoryService {
	genService := service.NewCRUDService(repo)
	return &categoryService{CRUDService: genService}
}
