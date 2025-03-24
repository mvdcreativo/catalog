package mql_request_filter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FilterFromRequest(c *gin.Context) (map[string]interface{}, error) {
	var filter map[string]interface{}

	switch c.Request.Method {
	case http.MethodGet:
		filterStr := c.Query("filter")
		if filterStr != "" {
			if err := json.Unmarshal([]byte(filterStr), &filter); err != nil {
				return nil, fmt.Errorf("filtro inválido (GET): %w", err)
			}
		}

	case http.MethodPost:
		if err := c.ShouldBindJSON(&filter); err != nil {
			return nil, fmt.Errorf("filtro inválido (POST): %w", err)
		}
	default:
		return nil, fmt.Errorf("método no soportado: %s", c.Request.Method)
	}

	// Si no se proporcionó ningún filtro, devolver mapa vacío
	if filter == nil {
		filter = map[string]interface{}{}
	}

	return filter, nil
}
