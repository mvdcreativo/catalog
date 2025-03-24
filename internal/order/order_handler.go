package order

import (
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/handler"
)


type OrderHandler struct {
	handler.CRUDHandler[Order]
}

func NewOrderHandler(service OrderService) *OrderHandler {
	genHandler := handler.NewCRUDHandler(service)
	return &OrderHandler{CRUDHandler: *genHandler}
}