package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/product"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/middleware"
)

// ProductRoutes registra las rutas de productos
func RegisterProductRoutes(rg *gin.RouterGroup, h *product.ProductHandler) {
	group := rg.Group("/products")
	{
		group.GET("/", h.FindAll)
		group.POST("", middleware.BindAndValidate[product.Product](), h.Insert)
		group.GET("/:id", middleware.ObjectIDMiddleware(), h.FindByID)
		group.PUT("/:id", middleware.BindAndValidate[product.Product](), middleware.ObjectIDMiddleware(), h.Update)
		group.DELETE("/:id", middleware.ObjectIDMiddleware(), h.Delete)
		group.POST("/upload_images/:id", middleware.ObjectIDMiddleware(), h.UploadImages)
		group.DELETE("/delete_images/:id", middleware.ObjectIDMiddleware(), h.DeleteImages)
	}

}
