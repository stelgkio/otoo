package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)


type AnalyticsBase struct {
	ID                primitive.ObjectID `bson:"_id,omitempty"`
	ProjectID         string             `bson:"projectId"`
	Timestamp 		  time.Time          `bson:"timestamp,omitempty"`
	TotalOrders       int64      		 `json:"total_orders"`
	TotalRevenue      float64    		 `json:"total_revenue"`				
	ActiveOrders      int64      		 `json:"active_orders"` 
	ActiveOrderRate   float64    		 `json:"active_order_rate"` 
	
}

// WeeklyAnalytics represents the analytics data for the last week.
type WeeklyAnalytics struct {
	AnalyticsBase
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
func  NewWeeklyAnalytics(projectID string, totalOrders,activeOrders int64,totalRevenue float64, startDate, endDate time.Time) WeeklyAnalytics {
	return WeeklyAnalytics{
		AnalyticsBase: AnalyticsBase{
			ProjectID: projectID,
			TotalOrders: totalOrders,
			TotalRevenue: totalRevenue,
			ActiveOrders: activeOrders,			
			
		},
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// MonthlyAnalytics represents the analytics data for the last month.
type MonthlyAnalytics struct {
	AnalyticsBase
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
func NewMonthlyAnalytics(projectID string, totalOrders,activeOrders int64,totalRevenue float64, startDate, endDate time.Time) MonthlyAnalytics {
	return MonthlyAnalytics{
		AnalyticsBase: AnalyticsBase{
			ProjectID: projectID,
			TotalOrders: totalOrders,
			TotalRevenue: totalRevenue,
			ActiveOrders: activeOrders,			
			
		},
		StartDate: startDate,
		EndDate:   endDate,
	}
}

// YearlyAnalytics represents the analytics data for the last year.
type YearlyAnalytics struct {
	AnalyticsBase
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
func  NewYearlyAnalytics(projectID string, totalOrders,activeOrders int64,totalRevenue float64, startDate, endDate time.Time) YearlyAnalytics {
	return YearlyAnalytics{
		AnalyticsBase: AnalyticsBase{
			ProjectID: projectID,
			TotalOrders: totalOrders,
			TotalRevenue: totalRevenue,
			ActiveOrders: activeOrders,	
			
		},
		StartDate: startDate,
		EndDate:   endDate,
	}
}
// calculatePercentages calculates and updates the percentage fields in AnalyticsBase.
func (a *AnalyticsBase) CalculatePercentages() {
	if a.TotalOrders > 0 {
		a.ActiveOrderRate = (float64(a.ActiveOrders) / float64(a.TotalOrders)) * 100
	} else {
		a.ActiveOrderRate = 0
	}
}


// ComparisonResult holds the results of comparing two periods.
type ComparisonResult struct {
	TotalOrdersChange       int64   `json:"total_orders_change"`
	TotalRevenueChange      float64 `json:"total_revenue_change"`
	ActiveOrdersChange      int64   `json:"active_orders_change"`
	ActiveOrderRateChange   float64 `json:"active_order_rate_change"`
}

// compareAnalytics compares the metrics between two WeeklyAnalytics periods.
func CompareAnalytics(current, previous AnalyticsBase) ComparisonResult {
	return ComparisonResult{
		TotalOrdersChange:       PercentageOrdersChange(current.TotalOrders, previous.TotalOrders),
		TotalRevenueChange:      PercentageBalanceChange(current.TotalRevenue, previous.TotalRevenue),
		ActiveOrdersChange:      PercentageOrdersChange(current.ActiveOrders, previous.ActiveOrders),
		ActiveOrderRateChange:   PercentageBalanceChange(current.ActiveOrderRate, previous.ActiveOrderRate),
	}
}

// percentageChange calculates the percentage change between two values.
func PercentageBalanceChange(current, previous float64) float64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return ((current - previous) / previous) * 100
}
// percentageChange calculates the percentage change between two values.
func PercentageOrdersChange(current, previous int64) int64 {
	if previous == 0 {
		if current == 0 {
			return 0
		}
		return 100
	}
	return ((current - previous) / previous) * 100
}
