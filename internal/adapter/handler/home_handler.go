package handler

import (
	"log/slog"

	"github.com/labstack/echo/v4"

	v "github.com/stelgkio/otoo/internal/adapter/web/view"
	con "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact"
	conf "github.com/stelgkio/otoo/internal/adapter/web/view/component/contact/dashboard-contact-form/contact-form"
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

// Post /dashboard/contact
func (h HomeHandler) DashboardContactForm(ctx echo.Context) error {
	req := new(domain.ContactRequest)
	if err := ctx.Bind(req); err != nil {
		slog.Error("Create contact binding error", "error", err)
		return r.Render(ctx, conf.ContactForm(true, false, nil, req)) // add errors
	}
	validationErrors := req.Validate()
	if len(validationErrors) > 0 {
		slog.Error("Create contact binding error", "error", validationErrors)
		return r.Render(ctx, conf.ContactForm(false, false, validationErrors, req)) // add errors

	}
	err := h.cont.InsertContact(ctx, req)
	if err != nil {
		slog.Error("Create contact binding error", "error", err)
		return r.Render(ctx, conf.ContactForm(true, false, nil, req)) // add errors
	}

	return r.Render(ctx, conf.ContactForm(false, true, nil, req)) // add notification message
}
