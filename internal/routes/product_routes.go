package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
)

// ProductRoutes registra las rutas de productos
func ProductRoutes(
	router *gin.RouterGroup,
	productHandler *handlers.ProductHandler) {
	products := router.Group("/products")
	{
		products.GET("/", productHandler.FindAll)
		products.POST("/", productHandler.Insert)
		products.GET("/:id", productHandler.FindByID)
		products.PUT("/:id", productHandler.Update)
		products.DELETE("/:id", productHandler.Delete)
	}
}
