package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/category"
)

// CategoryRoutes registra las rutas relacionadas con categorías.
func RegisterCategoryRoutes(rg *gin.RouterGroup, h *category.CategoryHandler) {
	group := rg.Group("/categories")
	group.GET("", h.FindAll)
	group.POST("", h.Insert)
	group.GET("/:id", h.FindByID)
	group.PUT("/:id", h.Update)
	group.DELETE("/:id", h.Delete)
}
