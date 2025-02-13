package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

// TokenService is an interface for interacting with token-related business logic
type TokenService interface {
	// CreateToken creates a new token for a given user
	CreateToken(user *domain.User) (string, error)
	// VerifyToken verifies the token and returns the payload
	//	VerifyToken(token string) (*domain.TokenPayload, error)
}

// UserService is an interface for interacting with user authentication-related business logic
type AuthService interface {
	// Login authenticates a user by email and password and returns a token
	Login(ctx echo.Context, email, password string) (string, error)
	Logout(ctx echo.Context) error
	ForgotPassword(ctx echo.Context, email string) error
	ResetPassword(ctx echo.Context) error
	ValidateCurrentPassword(ctx echo.Context, password, email string) (bool, error)
}
