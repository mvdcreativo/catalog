package responses

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RespondSuccess envía una respuesta exitosa estandarizada.
func RespondSuccess(c *gin.Context, status int, message string, data ...interface{}) {
	var responseData interface{}
	if len(data) > 0 {
		responseData = data[0]
	} else {
		responseData = nil
	}

	c.JSON(status, StandardResponse{
		Success: true,
		Message: message,
		Data:    responseData,
	})
}

// RespondPaginated envía una respuesta estandarizada con paginación.
func RespondPaginated(c *gin.Context, status int, message string, data interface{}, page int, limit int, total int64) {
	c.JSON(status, PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Page:    page,
		Limit:   limit,
		Total:   total,
	})
}

// RespondError envía una respuesta de error estandarizada.
func RespondError(c *gin.Context, status int, message interface{}) {
	switch msg := message.(type) {
	case string:
		c.JSON(status, ErrorResponse{
			Success: false,
			Error:   msg,
			Code:    status,
		})
	case error:
		c.JSON(status, ErrorResponse{
			Success: false,
			Error:   msg.Error(),
			Code:    status,
		})
	case map[string]string:
		// Podés usar otro struct para esto si querés
		c.JSON(status, gin.H{
			"success": false,
			"errors":  msg,
			"code":    status,
		})
	default:
		c.JSON(status, ErrorResponse{
			Success: false,
			Error:   fmt.Sprintf("Unexpected error: %v", msg),
			Code:    status,
		})
	}
}
