package woocommerce

import (
	"github.com/labstack/echo/v4"
	"github.com/stelgkio/otoo/internal/core/port"
)

// ReportService represents the woocommerce report service
type ReportService struct {
	ps port.ProjectService
}

// NewWoocommerceReportService creates a new woocommerce report service instance
func NewWoocommerceReportService(ps port.ProjectService) *ReportService {
	return &ReportService{
		ps,
	}
}

// GetCustomerTotalCount get customer report total count
func (s *ReportService) GetCustomerTotalCount(ctx echo.Context, projectID string) (int, error) {
	project, err := s.ps.GetProjectByID(ctx, projectID)
	if err != nil {
		return 0, err
	}

	client := InitClient(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)
	customerReportTotalCount, err := client.Report.GetTotalCustomers(nil)
	total := 0
	for _, customerCount := range customerReportTotalCount {
		total += customerCount.Total
	}

	if err != nil {
		return 0, err
	}

	return total, nil
}

// GetOrderTotalCount get order report total count
func (s *ReportService) GetOrderTotalCount(ctx echo.Context, projectID string) (int, error) {
	project, err := s.ps.GetProjectByID(ctx, projectID)
	if err != nil {
		return 0, err
	}

	client := InitClient(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)
	orderTotalCount, err := client.Report.GetTotalOrders(nil)
	total := 0
	for _, orderCount := range orderTotalCount {
		total += orderCount.Total
	}
	if err != nil {
		return 0, err
	}
	return total, nil
}

// GetProductTotalCount get product report total count
func (s *ReportService) GetProductTotalCount(ctx echo.Context, projectID string) (int, error) {
	project, err := s.ps.GetProjectByID(ctx, projectID)
	if err != nil {
		return 0, err
	}

	client := InitClient(project.WoocommerceProject.ConsumerKey, project.WoocommerceProject.ConsumerSecret, project.WoocommerceProject.Domain)
	productReportTotalCount, err := client.Report.GetTotalProducts(nil)
	total := 0
	for _, productCount := range productReportTotalCount {
		total += productCount.Total
	}
	if err != nil {
		return 0, err
	}
	return total, nil
}

// GetCustomerTotalCountTestCredential get customer report total count
func (s *ReportService) GetCustomerTotalCountTestCredential(ctx echo.Context, customerKey string, customerSecret string, domainURL string) (int, error) {

	client := InitClient(customerKey, customerSecret, domainURL)
	customerReportTotalCount, err := client.Report.GetTotalCustomers(nil)
	if err != nil {
		return 0, err
	}
	total := 0
	for _, customerCount := range customerReportTotalCount {
		total += customerCount.Total
	}

	return total, nil
}
