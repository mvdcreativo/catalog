package handlers

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/models"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/services"
)

// ProductHandler encapsula el servicio de Product.
type ProductHandler struct {
	CRUDHandler[models.Product]
}

// NewProductHandler crea una nueva instancia de ProductHandler con el servicio inyectado.
func NewProductHandler(service services.ProductService) *ProductHandler {
	genHandler := NewCRUDHandler(service)
	return &ProductHandler{CRUDHandler: *genHandler}
}
