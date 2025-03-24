package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler handles the health check requests.
func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
