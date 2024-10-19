package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AnalyticsBase struct {
	ID              primitive.ObjectID `bson:"_id,omitempty"`
	TotalOrders     int64              `json:"total_orders" bson:"total_orders"`
	TotalRevenue    float64            `json:"total_revenue" bson:"total_revenue"`
	ActiveOrders    int64              `json:"active_orders" bson:"active_orders"`
	ActiveOrderRate float64            `json:"order_rate" bson:"order_rate"`
}

// WeeklyAnalytics represents the analytics data for the last week.
type WeeklyAnalytics struct {
	AnalyticsBase
	ComparisonResult
	ProjectID string    `bson:"projectId" json:"project_id"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp"`
	StartDate time.Time `json:"start_date" bson:"start_date"`
	EndDate   time.Time `json:"end_date" bson:"end_date"`
}

func NewWeeklyAnalytics(projectID string, totalOrders, activeOrders int64, totalRevenue float64, startDate, endDate time.Time) *WeeklyAnalytics {
	return &WeeklyAnalytics{
		AnalyticsBase: AnalyticsBase{

			TotalOrders:  totalOrders,
			TotalRevenue: totalRevenue,
			ActiveOrders: activeOrders,
		},
		ProjectID: projectID,
		Timestamp: time.Now().UTC(),
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// AddComparisonResult add
func (w *WeeklyAnalytics) AddComparisonResult(result ComparisonResult) {
	w.ComparisonResult = result
}

// MonthlyOrderCountAnalytics represents the analytics data for the last month.
type MonthlyOrderCountAnalytics struct {
	MonthyData map[string]int `bson:"monthlydata" json:"monthlydata"`
	ProjectID  string         `bson:"projectId" json:"project_id"`
	Timestamp  time.Time      `bson:"timestamp,omitempty" json:"timestamp"`
	StartDate  time.Time      `json:"start_date"`
	EndDate    time.Time      `json:"end_date"`
}

// NewMonthlyAnalytics creates
func NewMonthlyAnalytics(projectID string, monthyData map[string]int, startDate, endDate time.Time) MonthlyOrderCountAnalytics {
	return MonthlyOrderCountAnalytics{
		MonthyData: monthyData,
		ProjectID:  projectID,
		StartDate:  startDate,
		EndDate:    endDate,
	}
}

// YearlyAnalytics represents the analytics data for the last year.
type YearlyAnalytics struct {
	AnalyticsBase
	ProjectID string    `bson:"projectId" json:"project_id"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

// NewYearlyAnalytics new
func NewYearlyAnalytics(projectID string, totalOrders, activeOrders int64, totalRevenue float64, startDate, endDate time.Time) YearlyAnalytics {
	return YearlyAnalytics{
		AnalyticsBase: AnalyticsBase{

			TotalOrders:  totalOrders,
			TotalRevenue: totalRevenue,
			ActiveOrders: activeOrders,
		},
		ProjectID: projectID,
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// CalculatePercentages calculates and updates the percentage fields in AnalyticsBase.
func (a *AnalyticsBase) CalculatePercentages() {
	if a.TotalOrders > 0 {
		a.ActiveOrderRate = (float64(a.ActiveOrders) / float64(a.TotalOrders)) * 100
	} else {
		a.ActiveOrderRate = 0
	}
}

// ComparisonResult holds the results of comparing two periods.
type ComparisonResult struct {
	TotalOrdersChange     int64   `json:"total_orders_change"`
	TotalRevenueChange    float64 `json:"total_revenue_change"`
	ActiveOrdersChange    int64   `json:"active_orders_change"`
	ActiveOrderRateChange float64 `json:"active_order_rate_change"`
}

// CompareAnalytics compares the metrics between two WeeklyAnalytics periods.
func CompareAnalytics(current, previous AnalyticsBase) ComparisonResult {
	return ComparisonResult{
		TotalOrdersChange:     PercentageOrdersChange(current.TotalOrders, previous.TotalOrders),
		TotalRevenueChange:    PercentageBalanceChange(current.TotalRevenue, previous.TotalRevenue),
		ActiveOrdersChange:    PercentageOrdersChange(current.ActiveOrders, previous.ActiveOrders),
		ActiveOrderRateChange: PercentageBalanceChange(current.ActiveOrderRate, previous.ActiveOrderRate),
	}
}

// PercentageBalanceChange calculates the percentage change between two values.
func PercentageBalanceChange(current, previous float64) float64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return ((current - previous) / previous) * 100
}

// PercentageOrdersChange calculates the percentage change between two values.
func PercentageOrdersChange(current, previous int64) int64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return ((current - previous) / previous) * 100
}
