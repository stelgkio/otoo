package handler

import (
	"github.com/labstack/echo/v4"
	e "github.com/stelgkio/otoo/internal/adapter/web/view/extention"
	"github.com/stelgkio/otoo/internal/core/util"
)

// Extention get extention
func (dh *DashboardHandler) Extention(ctx echo.Context) error {

	return util.Render(ctx, e.Extentions())
}

// AcsCourier get extention courier page
func (dh *DashboardHandler) AcsCourier(ctx echo.Context) error {

	return util.Render(ctx, e.ASC_Courier())
}

// WalletExpenses get extention
func (dh *DashboardHandler) WalletExpenses(ctx echo.Context) error {

	return util.Render(ctx, e.WalletExpenses())
}

// DataSynchronizer get extention
func (dh *DashboardHandler) DataSynchronizer(ctx echo.Context) error {

	return util.Render(ctx, e.DataSynchronizer())
}
