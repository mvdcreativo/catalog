package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/product"
)

// SetupRoutes define las rutas del API
func SetupRoutes(r *gin.Engine,
	ph *product.ProductHandler,
	ch *category.CategoryHandler,
) {

	api := r.Group("/api/v1")
	RegisterProductRoutes(api, ph)
	RegisterCategoryRoutes(api, ch)

	api.GET("/health_check", HealthHandler)
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
