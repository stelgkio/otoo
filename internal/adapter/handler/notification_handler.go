package handler

import (
	"github.com/labstack/echo/v4"
	t "github.com/stelgkio/otoo/internal/adapter/web/view/component/navigation/notification"
	"github.com/stelgkio/otoo/internal/core/util"
)

// FindNotification find all notification by projectID
func (dh *DashboardHandler) FindNotification(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	notifications, err := dh.notificationSvc.FindNotification(projectID, 10, 1, "timestamp", "", true)
	if err != nil {
		return err
	}

	return util.Render(ctx, t.NotificationIcon(notifications, projectID))
}

// DeleteNotification  delete a  notification by Id
func (dh *DashboardHandler) DeleteNotification(ctx echo.Context) error {
	projectID := ctx.Param("projectId")
	notificationID := ctx.Param("notifiactionId")
	err := dh.notificationSvc.DeleteNotification(ctx, projectID, notificationID)
	if err != nil {
		return err
	}
	notifications, err := dh.notificationSvc.FindNotification(projectID, 10, 1, "timestamp", "", true)
	if err != nil {
		return err
	}

	return util.Render(ctx, t.NotificationIcon(notifications, projectID))
}

// DeleteAllNotification  delete all  notification by Id
func (dh *DashboardHandler) DeleteAllNotification(ctx echo.Context) error {
	projectID := ctx.Param("projectId")

	notifications, err := dh.notificationSvc.FindNotification(projectID, 1000, 1, "timestamp", "", true)
	if err != nil {
		return err
	}
	for _, item := range notifications {
		notificationID := item.ID
		err := dh.notificationSvc.DeleteNotification(ctx, projectID, notificationID.Hex())
		if err != nil {
			return err
		}
	}
	notifications2, err := dh.notificationSvc.FindNotification(projectID, 1000, 1, "timestamp", "", true)
	if err != nil {
		return err
	}
	return util.Render(ctx, t.NotificationIcon(notifications2, projectID))
}
