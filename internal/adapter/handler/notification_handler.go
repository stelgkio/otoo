package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// FindNotification find all notification by projectID
func (dh *DashboardHandler) FindNotification(ctx echo.Context) error {
	//	projectID := ctx.Param("projectId")
	return ctx.NoContent(http.StatusOK)
}
