package models

// ErrorResponse represents the structure of error responses
type ErrorResponse struct {
    Error string `json:"error"`
}

// MessageResponse represents a simple message response
type MessageResponse struct {
    Message string `json:"message"`
}