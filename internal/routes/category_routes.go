package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/handlers"
)

// CategoryRoutes registra las rutas relacionadas con categorías.
func CategoryRoutes(
	api *gin.RouterGroup,
	categoryHandler *handlers.CategoryHandler,
) {
	categoryGroup := api.Group("/categories")
	{
		categoryGroup.GET("/", categoryHandler.FindAll)
		categoryGroup.GET("/:id", categoryHandler.FindByID)
		categoryGroup.POST("/", categoryHandler.Insert)
		categoryGroup.PUT("/:id", categoryHandler.Update)
		categoryGroup.DELETE("/:id", categoryHandler.Delete)
	}
}
