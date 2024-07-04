package service

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/auth"
	"github.com/stelgkio/otoo/internal/core/port"
	e "github.com/stelgkio/otoo/internal/core/util"
)

/**
 * AuthService implements port.AuthService interface
 * and provides an access to the user repository
 * and token service
 */
type AuthService struct {
	repo port.UserRepository
	ts   port.TokenService
}

// NewAuthService creates a new auth service instance
func NewAuthService(repo port.UserRepository, ts port.TokenService) *AuthService {
	return &AuthService{
		repo,
		ts,
	}
}

// Login gives a registered user an access token if the credentials are valid
func (as *AuthService) Login(ctx echo.Context, email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		if err == e.ErrDataNotFound {
			return "", e.ErrInvalidCredentials
		}
		return "", e.ErrInternal
	}

	err = e.ComparePassword(password, user.Password)
	if err != nil {
		return "", e.ErrInvalidCredentials
	}

	err = auth.GenerateTokensAndSetCookies(user, ctx)
	if err != nil {
		return "", e.ErrTokenCreation
	}

	return "", nil
}

func (as *AuthService) Logout(ctx echo.Context) error {

	userId, err := auth.GetUserId(ctx)
	if err != nil {
		return e.ErrInternal
	}
	user, err := as.repo.GetUserById(ctx, userId)
	if err != nil {
		if err == e.ErrDataNotFound {
			return e.ErrInvalidCredentials
		}
		return e.ErrInternal
	}
	err = auth.RemoveTokensAndDeleteCookies(user, ctx)
	if err != nil {
		return e.ErrTokenCreation
	}

	return nil
}

func (as *AuthService) ForgotPassword(ctx echo.Context, email string) error {
	user, err := as.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}
	link, _ := e.ResetPasswordLinkGenerator(email)
	if user != nil && link != "" {
		NewSmtpService().SendForgetPasswordEmail(ctx, email, user.Name, user.LastName, link)
	}

	return nil
}

func (as *AuthService) ResetPassword(ctx echo.Context) error {
	return nil
}
