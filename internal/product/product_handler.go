package product

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/handler"
)

// ProductHandler encapsula el servicio de Product.
type ProductHandler struct {
	handler.CRUDHandler[Product]
}

// NewProductHandler crea una nueva instancia de ProductHandler con el servicio inyectado.
func NewProductHandler(service ProductService) *ProductHandler {
	genHandler := handler.NewCRUDHandler(service)
	return &ProductHandler{CRUDHandler: *genHandler}
}
