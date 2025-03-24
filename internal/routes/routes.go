package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
)

// SetupRoutes define las rutas del API
func SetupRoutes(r *gin.Engine,
	ph *handlers.ProductHandler,
	ch *handlers.CategoryHandler,
) {
	api := r.Group("/api")
	RegisterProductRoutes(api, ph)
	RegisterCategoryRoutes(api, ch)
}
