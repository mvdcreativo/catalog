package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/product"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes define las rutas del API
func SetupRoutes(r *gin.Engine,
	ph *product.ProductHandler,
	ch *category.CategoryHandler,
) {

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	api := r.Group("/api/v1")
	RegisterProductRoutes(api, ph)
	RegisterCategoryRoutes(api, ch)

	api.GET("/health_check", HealthHandler)
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
