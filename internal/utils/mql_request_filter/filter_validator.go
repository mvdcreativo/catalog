package mql_request_filter

import (
	"errors"
)

// Operadores permitidos
var allowedOperators = map[string]bool{
	"$eq":  true,
	"$ne":  true,
	"$gt":  true,
	"$gte": true,
	"$lt":  true,
	"$lte": true,
	"$in":  true,
	"$nin": true,
	// Añadir más operadores según sea necesario
}

// Validar y sanitizar filtros
func ValidateAndSanitizeFilter(filter map[string]interface{}, whitelist map[string]bool) (map[string]interface{}, error) {
	sanitizedFilter := make(map[string]interface{})

	for key, value := range filter {
		// Verificar si el campo está en la lista blanca
		if !whitelist[key] {
			return nil, errors.New("campo no permitido en el filtro: " + key)
		}

		// Validar operadores si el valor es un mapa
		if valueMap, ok := value.(map[string]interface{}); ok {
			sanitizedValue := make(map[string]interface{})
			for op, v := range valueMap {
				if !allowedOperators[op] {
					return nil, errors.New("operador no permitido: " + op)
				}
				sanitizedValue[op] = v
			}
			sanitizedFilter[key] = sanitizedValue
		} else {
			// Si no es un mapa, asignar el valor directamente
			sanitizedFilter[key] = value
		}
	}

	return sanitizedFilter, nil
}
