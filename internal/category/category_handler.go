package category

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/handler"
)

type CategoryHandler struct {
	handler.CRUDHandler[Category]
}

// NewCategoryHandler crea una nueva instancia de CategoryHandler con el servicio inyectado.
func NewCategoryHandler(service CategoryService) *CategoryHandler {
	genHandler := handler.NewCRUDHandler(service)
	return &CategoryHandler{CRUDHandler: *genHandler}
}
