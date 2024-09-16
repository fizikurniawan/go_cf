package models

// General success response
type SuccessResponse struct {
	Status     string      `json:"status"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Pagination struct
type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Total   int `json:"total"`
}

// Error response
type ErrorResponse struct {
	Status  string              `json:"status"`
	Message string              `json:"message"`
	Errors  map[string][]string `json:"errors,omitempty"`
}
