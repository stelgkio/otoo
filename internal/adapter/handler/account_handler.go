package handler

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	l "github.com/stelgkio/otoo/internal/adapter/web/view/account/login"
	reg "github.com/stelgkio/otoo/internal/adapter/web/view/account/register"
	"github.com/stelgkio/otoo/internal/core/domain"
	"github.com/stelgkio/otoo/internal/core/port"
	r "github.com/stelgkio/otoo/internal/core/util"
)

// AuthHandler represents the HTTP handler for authentication-related requests
type AuthHandler struct {
	svc port.AuthService
	urs port.UserService
}

// NewAuthHandler creates a new AuthHandler instance
func NewAuthHandler(svc port.AuthService, urs port.UserService) *AuthHandler {
	return &AuthHandler{
		svc,
		urs,
	}
}

// authResponse represents an authentication response body
type authResponse struct {
	AccessToken string `json:"token"`
}

func AuthResponse(token string) authResponse {
	return authResponse{
		AccessToken: token,
	}
}

// loginRequest represents the request body for logging in a user
type loginRequest struct {
	Email    string `form:"email" validate:"required,email" example:"test@example.com"`
	Password string `form:"password" validate:"required,min=8" example:"12345678" minLength:"8"`
}

// Login godoc
//
//	@Summary		Login and get an access token
//	@Description	Logs in a registered user and returns an access token if the credentials are valid.
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			request	body		loginRequest	true	"Login request body"
//	@Success		200		{object}	authResponse	"Succesfully logged in"
//	@Failure		400		{object}	errorResponse	"Validation error"
//	@Failure		401		{object}	errorResponse	"Unauthorized error"
//	@Failure		500		{object}	errorResponse	"Internal server error"
//	@Router			/login [post]
func (ah *AuthHandler) Login(ctx echo.Context) (err error) {

	req := new(loginRequest)

	if err := ctx.Bind(req); err != nil {
		return ctx.String(http.StatusBadRequest, "bad request")
	}

	_, err = ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		return r.Render(ctx, l.Login(err))
		//ctx.String(http.StatusBadRequest, err.Error())
	}

	//AuthResponse(token)

	return ctx.Redirect(http.StatusFound, "/dashboard")
}

// @Router			/login [get]
func (ah *AuthHandler) LoginForm(c echo.Context) error {
	return r.Render(c, l.Login(nil))

}

// registerRequest represents the request body for creating a user
type registerRequest struct {
	Email                string `form:"email" validate:"required,email"`
	Password             string `form:"password" validate:"required,min=8"`
	ConfirmationPassword string `form:"confirmationpassword" validate:"required,min=8"`
	Name                 string `form:"name" validate:"required"`
	LastName             string `form:"last_name" validate:"required"`
}

// @Router			/register [get]
func (ah *AuthHandler) RegisterForm(ctx echo.Context) error {
	return r.Render(ctx, reg.Register(0, nil, nil))
}

// @Router			/register [post]
func (ah *AuthHandler) Register(ctx echo.Context) error {

	req := new(registerRequest)
	if err := ctx.Bind(req); err != nil {
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
		//return ctx.String(http.StatusBadRequest, "bad request")
	}
	ctx.Validate(req)
	// validate email is not taken
	validate := validator.New()
	// validate password is the same as confirm password
	if req.Password != req.ConfirmationPassword {
		return r.Render(ctx, reg.Register(0, nil, fmt.Errorf("Invalid confirmation password")))
	}
	// Validate the User struct
	err := validate.Struct(req)
	if err != nil {
		// Validation failed, handle the error
		errors := err.(validator.ValidationErrors)
		return r.Render(ctx, reg.Register(0, errors, nil))
	}

	user, err := domain.NewUser(req.Email, req.Password, req.Name, req.LastName)
	if err != nil {
		slog.Error("error new user:", "StatusBadRequest", err)
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
	}

	_, err = ah.urs.CreateUser(ctx, user)
	if err != nil {
		slog.Error("error create new user:", "StatusBadRequest", err)
		return r.Render(ctx, reg.Register(http.StatusBadRequest, nil, nil))
	}

	_, err = ah.svc.Login(ctx, req.Email, req.Password)
	if err != nil {
		slog.Error("error new user create token:", "StatusInternalServerError", err)
		return r.Render(ctx, reg.Register(http.StatusInternalServerError, nil, nil))
	}
	return ctx.Redirect(http.StatusMovedPermanently, ctx.Echo().Reverse("dashboard"))

}
