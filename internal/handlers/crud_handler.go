package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/responses"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/services"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CRUDHandler[T services.EntityModel] struct {
	productService services.CRUDService[T]
}

// NewProductHandler crea una nueva instancia de CRUDHandler con el servicio inyectado.
func NewCRUDHandler[T services.EntityModel](productService services.CRUDService[T]) *CRUDHandler[T] {
	return &CRUDHandler[T]{
		productService: productService,
	}
}

// GetProducts maneja GET /products
func (h *CRUDHandler[T]) FindAll(c *gin.Context) {
	ctx := c.Request.Context()

	filter, err := mql_request_filter.FilterFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, limit := mql_request_filter.GetPaginationParams(c)

	products, total, err := h.productService.FindAll(ctx, filter, page, limit)
	if err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondPaginated(c, http.StatusOK, "Products list", products, page, limit, total)
}

// GetProductByID maneja GET /products/:id
func (h *CRUDHandler[T]) FindByID(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	product, err := h.productService.FindByID(ctx, id)
	if err != nil {
		responses.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	responses.RespondSuccess(c, http.StatusOK, "Successfully obtained", product)
}

// CreateProduct maneja POST /products
func (h *CRUDHandler[T]) Insert(c *gin.Context) {
	ctx := c.Request.Context()
	log.Println(ctx)

	var newItem T
	if err := c.ShouldBindJSON(&newItem); err != nil {
		responses.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Casting dinámico al tipo Trackable
	if item, ok := any(&newItem).(services.Trackable); ok {
		item.SetID(primitive.NewObjectID())
		item.SetCreationDate(time.Now())
		item.SetUpdateDate(time.Now())
	}

	if err := h.productService.Insert(ctx, &newItem); err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondSuccess(c, http.StatusCreated, "Created", newItem)
}

// UpdateProduct maneja PUT /products/:id
func (h *CRUDHandler[T]) Update(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		responses.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}
	// Verificar si el producto existe por su ID
	if _, err := h.productService.FindByID(ctx, id); err != nil {
		responses.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	var updates T
	if err := c.ShouldBindJSON(&updates); err != nil {
		responses.RespondError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Se asigna el ID del parámetro a la entidad y se actualiza la fecha de modificación
	// Casting dinámico al tipo Trackable
	if item, ok := any(&updates).(services.Trackable); ok {
		item.SetID(primitive.NewObjectID())
		item.SetCreationDate(time.Now())
		item.SetUpdateDate(time.Now())
	}

	// Se llama al servicio para actualizar el producto
	if err := h.productService.Update(ctx, id, &updates); err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	updatedProduct, err := h.productService.FindByID(ctx, id)
	if err != nil {
		responses.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	responses.RespondSuccess(c, http.StatusOK, "Updated", updatedProduct)
}

// DeleteProduct maneja DELETE /products/:id
func (h *CRUDHandler[T]) Delete(c *gin.Context) {
	ctx := c.Request.Context()
	log.Println(ctx)
	idParam := c.Param("id")
	// Verificar si el producto existe por su ID
	if _, err := h.productService.FindByID(ctx, idParam); err != nil {
		responses.RespondError(c, http.StatusNotFound, err.Error())
		return
	}
	// Se llama al servicio para eliminar el producto
	if err := h.productService.Delete(ctx, idParam); err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "")
}
