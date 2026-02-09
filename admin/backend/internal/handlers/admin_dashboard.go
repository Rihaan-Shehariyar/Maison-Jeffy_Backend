package handlers

import (
	"backend/admin/backend/internal/database"
	"backend/internal/auth/entity"
	order_entity "backend/internal/orders/entity"
	"time"

	"github.com/gin-gonic/gin"
)

func GetDashboard(c *gin.Context) {

	var totalOrders int64
	var totalUsers int64
	var totalRevenue int64

	database.DB.Model(&order_entity.Order{}).Count(&totalOrders)
	database.DB.Model(&entity.User{}).Count(&totalUsers)

	database.DB.Model(&order_entity.Order{}).Select("COALESCE(SUM(total_amount),0)").Scan(&totalRevenue)

	// Orders per day
	type DailyStat struct {
		Date  string  `json:"date"`
		Count int     `json:"count"`
		Sum   float64 `json:"sum"`
	}

	var dailyStats []DailyStat

	database.DB.Raw(`
		SELECT 
			DATE(created_at) as date,
			COUNT(*) as count,
			SUM(total_amount) as sum
		FROM orders
		WHERE created_at >= ?
		GROUP BY DATE(created_at)
		ORDER BY date
	`, time.Now().AddDate(0, 0, -7)).Scan(&dailyStats)

	c.JSON(200, gin.H{
		"total_orders":  totalOrders,
		"total_users":   totalUsers,
		"total_revenue": totalRevenue,
		"daily_stats":   dailyStats,
	})

}
