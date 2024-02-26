package handler

import "github.com/stelgkio/otoo/internal/core/port"

// UserHandler represents the HTTP handler for user-related requests
type UserHandler struct {
	svc port.UserService
}

// NewUserHandler creates a new UserHandler instance
func NewUserHandler(svc port.UserService) *UserHandler {
	return &UserHandler{
		svc,
	}
}
