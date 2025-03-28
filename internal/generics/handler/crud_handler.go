package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/generics/service"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/interfaces/i_crud"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/responses"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/mql_request_filter"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CRUDHandler[T mql_request_filter.EntityModel] struct {
	curudService i_crud.CRUDService[T]
}

func NewCRUDHandler[T mql_request_filter.EntityModel](crudService i_crud.CRUDService[T]) *CRUDHandler[T] {
	return &CRUDHandler[T]{
		curudService: crudService,
	}
}

func (h *CRUDHandler[T]) FindAll(c *gin.Context) {
	ctx := c.Request.Context()

	filter, err := mql_request_filter.FilterFromRequest(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	page, limit := mql_request_filter.GetPaginationParams(c)

	results, total, err := h.curudService.FindAll(ctx, filter, page, limit)
	if err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondPaginated(c, http.StatusOK, "Results list", results, page, limit, total)
}

func (h *CRUDHandler[T]) FindByID(c *gin.Context) {
	ctx := c.Request.Context()

	// Ya parseado por el middleware, lo recuperamos del contexto
	objID := c.MustGet("objectID").(primitive.ObjectID)
	id := objID.Hex() // Esto espera string
	entity, err := h.curudService.FindByID(ctx, id)
	// Valida existencia
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			responses.RespondError(c, http.StatusNotFound, "Item not found")
		} else {
			responses.RespondError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	responses.RespondSuccess(c, http.StatusOK, "Successfully obtained", entity)
}

func (h *CRUDHandler[T]) Insert(c *gin.Context) {
	ctx := c.Request.Context()

	raw, _ := c.Get("validatedRequest")
	newItem := raw.(T)

	// Casting dinámico al tipo Trackable
	if item, ok := any(&newItem).(service.Trackable); ok {
		item.SetID(primitive.NewObjectID())
		item.SetCreationDate(time.Now())
		item.SetUpdateDate(time.Now())
	}

	if err := h.curudService.Insert(ctx, &newItem); err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	responses.RespondSuccess(c, http.StatusCreated, "Created", newItem)
}

func (h *CRUDHandler[T]) Update(c *gin.Context) {
	ctx := c.Request.Context()

	// Ya parseado por el middleware, lo recuperamos del contexto
	objID := c.MustGet("objectID").(primitive.ObjectID)
	id := objID.Hex() // Esto espera string

	// Valida existencia
	if _, err := h.curudService.FindByID(ctx, id); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			responses.RespondError(c, http.StatusNotFound, "Item not found")
		} else {
			responses.RespondError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}

	raw, _ := c.Get("validatedRequest")
	updates := raw.(T)

	if item, ok := any(&updates).(service.Trackable); ok {
		item.SetUpdateDate(time.Now())
	}

	if err := h.curudService.Update(ctx, id, &updates); err != nil {
		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	updatedEntity, err := h.curudService.FindByID(ctx, id)
	if err != nil {
		responses.RespondError(c, http.StatusNotFound, err.Error())
		return
	}

	responses.RespondSuccess(c, http.StatusOK, "Updated", updatedEntity)
}

func (h *CRUDHandler[T]) Delete(c *gin.Context) {
	ctx := c.Request.Context()

	// Ya parseado por el middleware, lo recuperamos del contexto
	objID := c.MustGet("objectID").(primitive.ObjectID)
	id := objID.Hex() // Esto espera string

	// Valida existencia
	if _, err := h.curudService.FindByID(ctx, id); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			responses.RespondError(c, http.StatusNotFound, "Item not found")
		} else {
			responses.RespondError(c, http.StatusInternalServerError, err.Error())
		}
		return
	}
	// Se llama al servicio para eliminar el entity
	if err := h.curudService.Delete(ctx, id); err != nil {
		log.Print("❌ Error validando existencia")

		responses.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, "")
}
