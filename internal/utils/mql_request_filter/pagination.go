// /internal/utils/pagination.go
package mql_request_filter

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// Obtener parámetros de paginación
func GetPaginationParams(c *gin.Context) (int, int) {
	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		limit = 10
	}

	return page, limit
}

// Calcular opciones de paginación para bases de datos que utilizan offset
func GetOffset(page, limit int) int {
	return (page - 1) * limit
}
