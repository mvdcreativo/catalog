package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
)

// SetupRoutes define las rutas del API
func SetupRoutes(
	router *gin.Engine,
	productHandler *handlers.ProductHandler,
	categoryHandler *handlers.CategoryHandler,
) {
	api := router.Group("/api") // Prefijo común para todas las rutas

	ProductRoutes(api, productHandler)
	CategoryRoutes(api, categoryHandler)
}
