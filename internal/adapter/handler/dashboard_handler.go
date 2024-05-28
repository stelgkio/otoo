package handler

import (
	"github.com/stelgkio/otoo/internal/core/port"
)

type DashboardHandler struct {
	projectSvc port.ProjectService
	userSvc    port.UserService
}

func NewDashboardHandler(projectSvc port.ProjectService, userSvc port.UserService) *DashboardHandler {
	return &DashboardHandler{
		projectSvc: projectSvc,
		userSvc:    userSvc,
	}
}
