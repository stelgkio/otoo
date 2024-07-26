package service

import "github.com/stelgkio/otoo/internal/core/port"

// AnalyticsService implements port.AnalyticsService interface
type AnalyticsService struct {
	mogorepo port.WoocommerceRepository
}

// NewAnalyticsService creates a new analytics service instance
func NewAnalyticsService(mogorepo port.WoocommerceRepository) *AnalyticsService {
	return &AnalyticsService{
		mogorepo,
	}
}
