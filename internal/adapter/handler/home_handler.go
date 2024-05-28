package handler

import (
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	l "github.com/stelgkio/otoo/internal/adapter/web/template/components/dashboard/projectlist"
	t "github.com/stelgkio/otoo/internal/adapter/web/template/components/home"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

type HomeHandler struct {
	svc port.ProjectService
}

func NewHomeHandler(svc port.ProjectService) *HomeHandler {
	return &HomeHandler{
		svc,
	}
}

func (h HomeHandler) Home(c echo.Context) error {
	projects, err := h.svc.FindProjects(c, &domain.FindProjectRequest{}, 1, 10)
	if err != nil {
		return err
	}
	return r.Render(c, t.HomeComponent(projects))
}

func (h HomeHandler) ProjectList(c echo.Context) error {
	//page := c.QueryParam("page")
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		slog.Error("Invalid page number", err)
		return err
	}
	projects, err := h.svc.FindProjects(c, &domain.FindProjectRequest{}, page, 10)
	if err != nil {
		return err
	}
	return r.Render(c, l.ProjectListComponent(projects, page))
}
