package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/domains/product"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/responses"
	"github.com/mvdcreativo/e-commerce-saas/catalog/internal/utils/validation"
)

var validate = validator.New()

func Init() {
	_ = validate.RegisterValidation("status", ValidateProductStatus)
}

func BindAndValidate[T any]() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req T

		if err := c.ShouldBindJSON(&req); err != nil {
			responses.RespondError(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		if err := validate.Struct(req); err != nil {
			errs := make(map[string]string)
			for _, e := range err.(validator.ValidationErrors) {
				errs[e.Field()] = e.ActualTag() // Podés agregar mensaje personalizado si querés
			}
			responses.RespondError(c, http.StatusUnprocessableEntity, validation.BuildValidationErrors(err))
			c.Abort()
			return
		}

		// Seteo la request parseada para que el handler la use
		c.Set("validatedRequest", req)
		c.Next()
	}
}

var statusProductSet map[string]struct{}

func init() {
	statusProductSet = make(map[string]struct{})
	for _, s := range product.ValidStatusesProduct() {
		statusProductSet[s] = struct{}{}
	}
}

// ValidateStatus revisa si el valor está en el set permitido
func ValidateProductStatus(fl validator.FieldLevel) bool {
	status := strings.ToUpper(fl.Field().String())
	_, ok := statusProductSet[status]
	return ok
}
