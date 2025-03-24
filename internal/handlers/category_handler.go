package handlers

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/services"
)

type CategoryHandler struct {
	CRUDHandler[models.Category]
}

// NewCategoryHandler crea una nueva instancia de CategoryHandler con el servicio inyectado.
func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	genHandler := NewCRUDHandler(service)
	return &CategoryHandler{CRUDHandler: *genHandler}
}
