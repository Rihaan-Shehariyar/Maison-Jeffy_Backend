package handlers

import (
	"backend/admin/backend/internal/database"
	"backend/internal/auth/entity"
	order_entity "backend/internal/orders/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {

	var totalOrders int64
	var pendingOrders int64
	var successfulOrders int64
	var totalUsers int64
	var totalRevenue float64

	// KPI COUNTS

	database.DB.Model(&order_entity.Order{}).Count(&totalOrders)

	database.DB.Model(&order_entity.Order{}).
		Where("status != ?", "delivered").
		Count(&pendingOrders)

	database.DB.Model(&order_entity.Order{}).
		Where("status = ?", "delivered").
		Count(&successfulOrders)

	database.DB.Model(&entity.User{}).Count(&totalUsers)

	database.DB.Model(&order_entity.Order{}).
		Where("status = ?", "delivered").
		Select("COALESCE(SUM(total_amount),0)").
		Scan(&totalRevenue)

	// LAST 7 DAYS CHART DATA

	type ChartRow struct {
		Date  time.Time
		Count int64
		Sum   float64
	}

	var orderRows []ChartRow
	var revenueRows []ChartRow

	// Orders per day

	database.DB.Raw(`
		SELECT DATE(created_at) as date, COUNT(*) as count
		FROM orders
		WHERE created_at >= NOW() - INTERVAL '7 days'
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at)
	`).Scan(&orderRows)

	// Revenue per day (only delivered)

	database.DB.Raw(`
		SELECT DATE(created_at) as date, COALESCE(SUM(total_amount),0) as sum
		FROM orders
		WHERE status='delivered'
		AND created_at >= NOW() - INTERVAL '7 DAY'
		GROUP BY DATE(created_at)
		ORDER BY DATE(created_at)
	`).Scan(&revenueRows)

	// FORMAT FOR FRONTEND

	labels := []string{}
	orderValues := []int64{}
	revenueValues := []float64{}

	for i := 6; i >= 0; i-- {
		day := time.Now().AddDate(0, 0, -i).Format("2006-01-02")
		labels = append(labels, day)

		var orderCount int64 = 0
		for _, r := range orderRows {
			if r.Date.Format("2006-01-02") == day {
				orderCount = r.Count
			}
		}
		orderValues = append(orderValues, orderCount)

		var revenue float64 = 0
		for _, r := range revenueRows {
			if r.Date.Format("2006-01-02") == day {
				revenue = r.Sum
			}
		}
		revenueValues = append(revenueValues, revenue)
	}

	// RESPONSE

	c.JSON(200, gin.H{
		"totalOrders":      totalOrders,
		"pendingOrders":    pendingOrders,
		"successfulOrders": successfulOrders,
		"totalUsers":       totalUsers,
		"totalRevenue":     totalRevenue,

		"ordersChart": gin.H{
			"labels": labels,
			"values": orderValues,
		},
		"revenueChart": gin.H{
			"labels": labels,
			"values": revenueValues,
		},
	})
}
