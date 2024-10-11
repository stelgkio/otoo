package port

import (
	d "github.com/stelgkio/otoo/internal/core/domain"
	domain "github.com/stelgkio/otoo/internal/core/domain/woocommerce"
)

// AnalyticsRepository defines the methods for interacting with the BestSellers repository
type AnalyticsRepository interface {
	FindBestSellers(projectID string, size, page int) ([]*domain.ProductBestSellerRecord, error)
	CreateBestSellers(projectID string, data []*domain.ProductBestSellerRecord) error
	DeleteBestSellers(projectID string) error

	FindWeeklyBalance(projectID string, size, page int) ([]*domain.WeeklyAnalytics, error)
	FindLatestWeeklyBalance(projectID string) (*domain.WeeklyAnalytics, error)
	CreateWeeklyBalance(projectID string, data *domain.WeeklyAnalytics) error
	DeleteWeeklyBalance(projectID string) error

	FindMonthlyCount(projectID string, size, page int) ([]*domain.MonthlyOrderCountAnalytics, error)
	FindLatestMonthlyCount(projectID string) (*domain.MonthlyOrderCountAnalytics, error)
	CreateMonthlyCount(projectID string, data *domain.MonthlyOrderCountAnalytics) error
	DeleteMonthlyCount(projectID string) error
}

// ProductBestSellers defines the methods for interacting with the ProductBestSellers repository
type ProductBestSellers interface {
	RunAProductBestSellerInitializerJob(projectID string) error
	RunAProductBestSellerDailyJob() error
}

// OrderAnalyticsCron defines the methods for interacting with the OrderAnalyticsCron repository
type OrderAnalyticsCron interface {
	RunOrderWeeklyBalanceJob() error
	RunOrderWeeklyBalanceInitializeJob(project *d.Project, user *d.User) error
	RunOrderMonthlyCountJob() error
	RunOrderMonthlyCountInitializeJob(project *d.Project, user *d.User) error
}
