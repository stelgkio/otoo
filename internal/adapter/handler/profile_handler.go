package handler

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	reg "github.com/stelgkio/otoo/internal/adapter/web/view/account/register"
	h "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile"
	p "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/profile_password"
	pe "github.com/stelgkio/otoo/internal/adapter/web/view/component/profile/update_profile_error"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	"github.com/stelgkio/otoo/internal/core/util"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// registerRequest represents the request body for creating a user
type updateProfileRequest struct {
	Email    string `form:"email" validate:"required,email"`
	Name     string `form:"name" validate:"required"`
	LastName string `form:"last_name" validate:"required"`
}

// ProfileHandler represents the HTTP handler for user-related requests
type ProfileHandler struct {
	svc  port.UserService
	psrv port.ProjectService
	asrc port.AuthService
}

// NewProfileHandler creates a new UserHandler instance
func NewProfileHandler(svc port.UserService, psvc port.ProjectService, asrc port.AuthService) *ProfileHandler {
	return &ProfileHandler{
		svc,
		psvc,
		asrc,
	}
}

// Profile  @Router			/profile [get]
func (ph *ProfileHandler) Profile(ctx echo.Context) error {
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	return r.Render(ctx, h.Profile(user))

}

// ProfileUpdate 	@Router			/profile/update [post]
func (ph *ProfileHandler) ProfileUpdate(ctx echo.Context) error {
	req := new(updateProfileRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, h.Profile(nil))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userID)
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

// ProfileDelete 	@Router			/profile/delete [post]
func (ph *ProfileHandler) ProfileDelete(ctx echo.Context) error {
	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	err = ph.asrc.Logout(ctx)
	err = ph.svc.DeleteUser(ctx, userID)

	if err != nil {
		ctx.Response().Header().Set("HX-Redirect", "/index")
		return ctx.Redirect(http.StatusAccepted, "/index")
	}

	ctx.Response().Header().Set("HX-Redirect", "/index")
	return ctx.NoContent(http.StatusOK)

}

// ProfilePassword 	@Router			/profile/password [get]
func (ph *ProfileHandler) ProfilePassword(ctx echo.Context) error {
	req := new(domain.UpdatePasswordRequest)
	return r.Render(ctx, p.ProfilePassword(nil, req, false))
}

// UpdatePassword 	@Router			/profile/password/update [post]
func (ph *ProfileHandler) UpdatePassword(ctx echo.Context) error {
	req := new(domain.UpdatePasswordRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))

	}
	validationErrors := req.Validate()

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	correctpass, err := ph.asrc.ValidateCurrentPassword(ctx, req.CurrentPassword, user.Email)
	if err != nil {
		validationErrors["currentPassword"] = "Current password is incorrect"
	}
	if !correctpass {
		validationErrors["currentPassword"] = "Current password is incorrect"
	}
	if len(validationErrors) > 0 {
		return r.Render(ctx, p.ProfilePassword(validationErrors, req, false))
	}

	var hash util.Hash
	user.Password, err = hash.Generate(req.Password)
	ph.svc.UpdateUser(ctx, user)
	return r.Render(ctx, p.ProfilePassword(validationErrors, req, true))
}

// ValidateCurrentPassword 	@Router			/profile/password/validate [post]
func (ph *ProfileHandler) ValidateCurrentPassword(ctx echo.Context) error {

	req := new(domain.UpdatePasswordRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))

	}
	validationErrors := req.Validate()

	userID, err := auth.GetUserID(ctx)
	if err != nil {
		return err
	}
	user, err := ph.svc.GetUserById(ctx, userID)
	if err != nil {
		return err
	}

	correctpass, err := ph.asrc.ValidateCurrentPassword(ctx, req.CurrentPassword, user.Email)
	if err != nil {
		validationErrors["currentPassword"] = "Current password is incorrect"
	}
	if !correctpass {
		validationErrors["currentPassword"] = "Current password is incorrect"
	}

	return r.Render(ctx, p.CurrentPasswordValidation(req.CurrentPassword, validationErrors))
}

// ValidateNewPassword 	@Router			/profile/password/validate [post]
func (ph *ProfileHandler) ValidateNewPassword(ctx echo.Context) error {

	req := new(domain.UpdatePasswordRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))

	}
	validationErrors := req.Validate()
	return r.Render(ctx, p.NewPasswordValidation(req.Password, req.ConfirmationPassword, validationErrors))
}
