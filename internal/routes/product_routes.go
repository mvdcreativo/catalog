package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
)

// ProductRoutes registra las rutas de productos
func RegisterProductRoutes(rg *gin.RouterGroup, h *handlers.ProductHandler) {
	group := rg.Group("/products")
	group.GET("", h.FindAll)
	group.POST("", h.Insert)
	group.GET("/:id", h.FindByID)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
