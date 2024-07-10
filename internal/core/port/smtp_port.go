package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

type SmtpService interface {
	// InsertWoocommerceOrder inserts a new order into the database
	SendEmail(ctx echo.Context, sender, email, template, subject string, isHtml bool) error
	SendForgetPasswordEmail(ctx echo.Context, email, firstName, lastName, resetPasswordLink string) error
	SendContactEmail(ctx echo.Context, req *domain.ContactRequest) error
}
