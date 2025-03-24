package responses

// StandardResponse es el formato para respuestas exitosas sin paginación.
type StandardResponse struct {
	Success bool        `json:"success"`           // Siempre true.
	Message string      `json:"message,omitempty"` // Mensaje opcional.
	Data    interface{} `json:"data,omitempty"`    // Datos de la respuesta.
}

// PaginatedResponse extiende StandardResponse para incluir datos de paginación.
type PaginatedResponse struct {
	Success bool        `json:"success"`           // Siempre true.
	Message string      `json:"message,omitempty"` // Mensaje opcional.
	Data    interface{} `json:"data,omitempty"`    // Lista de items.
	Page    int         `json:"page"`              // Página actual.
	Limit   int         `json:"limit"`             // Límite de items por página.
	Total   int64       `json:"total"`             // Total de items encontrados.
}

// ErrorResponse es el formato para respuestas de error.
type ErrorResponse struct {
	Success bool   `json:"success"` // Siempre false.
	Error   string `json:"error"`   // Mensaje de error.
	Code    int    `json:"code"`    // Código HTTP de error.
}
