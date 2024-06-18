package port

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/domain"
)

type ContactRepository interface {
	// InsertWoocommerceOrder inserts a new order into the database
	InsertContact(ctx echo.Context, c *domain.Contact) error
}

type ContactService interface {
	// InsertWoocommerceOrder inserts a new order into the database
	InsertContact(ctx echo.Context, c *domain.ContactRequest) error
}
