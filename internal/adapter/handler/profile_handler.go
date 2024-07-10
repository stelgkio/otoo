package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	h "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/profile_password"
	pe "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/update_profile_error"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// registerRequest represents the request body for creating a user
type updateProfileRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Name     string `form:"name" validate:"required"`
	LastName string `form:"last_name" validate:"required"`
}

// UserHandler represents the HTTP handler for user-related requests
type ProfileHandler struct {
	svc  port.UserService
	psrv port.ProjectService
	asrc port.AuthService
}

// NewUserHandler creates a new UserHandler instance
func NewProfileHandler(svc port.UserService, psvc port.ProjectService, asrc port.AuthService) *ProfileHandler {
	return &ProfileHandler{
		svc,
		psvc,
		asrc,
	}
}

// @Router			/profile [get]
func (ph *ProfileHandler) Profile(ctx echo.Context) error {
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	return r.Render(ctx, h.Profile(user))

}

func (ph *ProfileHandler) ProfilePassword(ctx echo.Context) error {

	return r.Render(ctx, p.ProfilePassword())
}

func (ph *ProfileHandler) ProfileUpdate(ctx echo.Context) error {
	req := new(updateProfileRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, h.Profile(nil))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userId)
	if err != nil {
		return err
	}

	if user.Email != req.Email {
		userByEmail, err := ph.svc.GetUserByEmail(ctx, req.Email)
		if err != nil {
			return err
		}
		if userByEmail != nil {
			slog.Error(" ProfileUpdate:", "user email already exist", userByEmail.Email)
			user.Email = req.Email
			return r.Render(ctx, pe.ProfileUpdateError(user))
		}
	}
	user.Email = req.Email
	user.LastName = req.LastName
	user.Name = req.Name

	newUser, err := ph.svc.UpdateUser(ctx, user)
	if err != nil {
		return err
	}
	return r.Render(ctx, h.Profile(newUser))

}

func (ph *ProfileHandler) ProfileDelete(ctx echo.Context) error {
	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return err
	}

	err = ph.psrv.SoftDeleteProjects(ctx, userId)
	if err != nil {
		return err
	}

	err = ph.asrc.Logout(ctx)
	if err != nil {
		ctx.Response().Header().Set("HX-Redirect", "/index")
		return ctx.Redirect(http.StatusAccepted, "/index")
	}
	err = ph.svc.DeleteUser(ctx, userId)
	if err != nil {
		return err
	}
	ctx.Response().Header().Set("HX-Redirect", "/index")
	return ctx.NoContent(http.StatusOK)

}
