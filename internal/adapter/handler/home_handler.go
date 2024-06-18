package handler

import (
	"log/slog"
	"strconv"

	"github.com/labstack/echo/v4"
	l "github.com/stelgkio/otoo/internal/adapter/web/template/components/dashboard/projectlist"
	t "github.com/stelgkio/otoo/internal/adapter/web/template/components/home"
	v "github.com/stelgkio/otoo/internal/adapter/web/view"
	con "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

type HomeHandler struct {
	svc  port.ProjectService
	cont port.ContactService
}

func NewHomeHandler(svc port.ProjectService, cont port.ContactService) *HomeHandler {
	return &HomeHandler{
		svc,
		cont,
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

// Post /contact
func (h HomeHandler) ContactForm(ctx echo.Context) error {
	req := new(domain.ContactRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create contact binding error", "error", err)
		return r.Render(ctx, con.ContantComponent()) // add errors
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		slog.Error("Create contact binding error", "error", validationErrors)
		return r.Render(ctx, con.ContantComponent()) // add errors

	}
	err := h.cont.InsertContact(ctx, req)
	if err != nil {
		slog.Error("Create contact binding error", "error", err)
		return r.Render(ctx, con.ContantComponent()) // add errors
	}

	return r.Render(ctx, v.IndexTemplate()) // add notification message
}
