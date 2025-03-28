package product

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/config"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/handler"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/storage"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/responses"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/file_validator"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/slices"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// ProductHandler encapsula el servicio de Product.
type ProductHandler struct {
	handler.CRUDHandler[Product]
	service ProductService
	cfg     config.Config
}

// NewProductHandler crea una nueva instancia de ProductHandler con el servicio inyectado.
func NewProductHandler(cfg *config.Config, service ProductService) *ProductHandler {
	genHandler := handler.NewCRUDHandler(service)
	return &ProductHandler{
		CRUDHandler: *genHandler,
		service:     service,
		cfg:         *cfg,
	}
}

// UploadImages sube las imagenes de un producto
func (h *ProductHandler) UploadImages(c *gin.Context) {
	ctx := c.Request.Context()
	productID := c.Param("id")
	if productID == "" {
		responses.RespondError(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	requestForm, err := c.MultipartForm()
	if err != nil {
		responses.RespondError(c, http.StatusBadRequest, "Invalid request")
		return
	}

	//validamos imagenes
	files := requestForm.File["images"]
	for _, f := range files {
		err := file_validator.ValidateFile(f, h.cfg.Upload.Images.MaxSizeMB, h.cfg.Upload.Images.AllowedTypes)
		if err != nil {
			log.Println("❌ Error validando imagen", f.Filename)
			responses.RespondError(c, http.StatusBadRequest, fmt.Sprintf("validation failed for %s: %v", f.Filename, err))
			return
		}
	}

	validFiles, err := h.service.UploadImages(ctx, requestForm, productID)
	if err != nil {
		responses.RespondError(c, http.StatusInternalServerError, "Upload failed")
		return
	}

	responses.RespondSuccess(c, http.StatusOK, "Successfully uploaded", gin.H{"files": validFiles})
}

func (h *ProductHandler) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Ya parseado por el middleware, lo recuperamos del contexto
	objID := c.MustGet("objectID").(primitive.ObjectID)
	id := objID.Hex() // Esto espera string

	// Valida existencia
	product, err := h.service.FindByID(ctx, id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			responses.RespondError(c, http.StatusNotFound, "Item not found")
		} else {
			responses.RespondError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	images := &product.Images
	imagesIds := slices.MapToProp(*images, func(f storage.FileObject) string {
		return f.ID
	})

	if err := h.service.DeleteImages(ctx, id, imagesIds); err != nil {
		log.Printf("❌ Error deleting images for product %s: %v", id, err)
		responses.RespondError(c, http.StatusInternalServerError, "Error deleting images")
		return
	}
	// Se llama al servicio para eliminar el entity
	if err := h.service.Delete(ctx, id); err != nil {
		log.Print("❌ Error validando existencia")
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("✅ Product %s deleted", id)
	c.JSON(http.StatusNoContent, "")
}

func (h *ProductHandler) DeleteImages(c *gin.Context) {
	ctx := c.Request.Context()

	var images []string
	if err := c.ShouldBindJSON(&images); err != nil {
		responses.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	productID := c.Param("id")
	if productID == "" {
		responses.RespondError(c, http.StatusBadRequest, "Product ID is required")
		return
	}

	product, err := h.service.FindByID(ctx, productID)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			responses.RespondError(c, http.StatusNotFound, "Item not found")
		} else {
			responses.RespondError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	updatedImages := slices.FilterByField(product.Images, images, func(f storage.FileObject) string {
		return f.ID
	})

	if err := h.service.DeleteImages(ctx, productID, images); err != nil {
		log.Printf("❌ Error deleting images for product %s: %v", productID, err)
		responses.RespondError(c, http.StatusInternalServerError, "Error deleting images")
		return
	}

	if err := h.service.Update(ctx, productID, &Product{Images: updatedImages}); err != nil {
		log.Printf("❌ Error updating product %s: %v", productID, err)
		responses.RespondError(c, http.StatusInternalServerError, "Error updating product")
		return
	}

	log.Printf("✅ Images deleted for product %s", productID)
	c.JSON(http.StatusNoContent, "")
}
