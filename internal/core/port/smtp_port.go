package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
	w "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
)

// SmtpService defines the methods for interacting with the SMTP service
type SmtpService interface {
	// InsertWoocommerceOrder inserts a new order into the database
	SendEmail(ctx echo.Context, sender, email, template, subject string, isHTML bool) error
	SendForgetPasswordEmail(ctx echo.Context, email, firstName, lastName, resetPasswordLink string) error
	SendContactEmail(ctx echo.Context, req *domain.ContactRequest) error
	SendWeeklyBalanceEmail(ctx echo.Context, req *w.WeeklyAnalytics, email, fullname string) error
	SendMonthlyOrdersEmail(ctx echo.Context, req *w.MonthlyOrderCountAnalytics, email, fullname string) error
}
