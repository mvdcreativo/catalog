package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/category"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/middleware"
)

func RegisterCategoryRoutes(rg *gin.RouterGroup, h *category.CategoryHandler) {
	group := rg.Group("/categories")
	group.GET("", h.FindAll)
	group.POST("", middleware.BindAndValidate[category.Category](), h.Insert)
	group.GET("/:id", h.FindByID)
	group.PUT("/:id", middleware.BindAndValidate[category.Category](), h.Update)
	group.DELETE("/:id", h.Delete)
}
