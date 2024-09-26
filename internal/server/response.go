package server

import (
	"time"
)

type APIResponse struct {
	Status      int       `json:"status"`
	Description string    `json:"description"`
	Data        any       `json:"data"`
	Time        time.Time `json:"time"`
	Message     string    `json:"message"`
}

func NewAPIResponse(status int, description string, data any, message string) *APIResponse {
	return &APIResponse{
		Status:      status,
		Description: description,
		Data:        data,
		Time:        time.Now().UTC(),
		Message:     message,
	}
}
