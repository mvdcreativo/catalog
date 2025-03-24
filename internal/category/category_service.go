package category

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
)

// CategoryService define las operaciones de negocio para Category.

type CategoryService interface {
	service.CRUDService[Category]
}

type categoryService struct {
	service.CRUDService[Category]
}

// NewCategoryService crea una nueva instancia de CategoryService inyectando el repositorio.
func NewCategoryService(repo CategoryRepository) CategoryService {
	genService := service.NewCRUDService(repo)
	return &categoryService{CRUDService: genService}
}
