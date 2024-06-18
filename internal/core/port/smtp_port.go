package port

import "github.com/labstack/echo/v4"

type SmtpService interface {
	// InsertWoocommerceOrder inserts a new order into the database
	SendEmail(ctx echo.Context) error
}
