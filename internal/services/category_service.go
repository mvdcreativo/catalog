package services

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/repositories"
)

// CategoryService define las operaciones de negocio para Category.

type CategoryService interface {
	CRUDService[models.Category]
}

type categoryService struct {
	CRUDService[models.Category]
}

// NewCategoryService crea una nueva instancia de CategoryService inyectando el repositorio.
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	genService := NewCRUDService(repo)
	return &categoryService{CRUDService: genService}
}
