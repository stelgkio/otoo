package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
)

// UserHandler represents the HTTP handler for user-related requests
type ProjectHandler struct {
	svc port.ProjectService
}

// NewProjectHandler creates a new ProjectHandler instance
func NewProjectHandler(svc port.ProjectService) *ProjectHandler {
	return &ProjectHandler{
		svc,
	}
}

func (ph *ProjectHandler) CreateProject(ctx echo.Context) error {
	req := new(domain.ProjectRequest)
	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	dom, err := ph.svc.CreateProject(ctx, req)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, dom)
}
