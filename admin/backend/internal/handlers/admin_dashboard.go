package handlers

import (
	"backend/admin/backend/internal/database"
	"backend/internal/auth/entity"
	order_entity "backend/internal/orders/entity"	

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {

	var totalOrders int64
	var pendingOrders int64
	var successfulOrders int64
	var totalUsers int64
	var totalRevenue float64

	// Orders count
	database.DB.Model(&order_entity.Order{}).Count(&totalOrders)

	// Pending orders
	database.DB.Model(&order_entity.Order{}).
		Where("status = ?", "pending").
		Count(&pendingOrders)

	// Successful orders
	database.DB.Model(&order_entity.Order{}).
		Where("status = ?", "delivered").
		Count(&successfulOrders)

	// Users count
	database.DB.Model(&entity.User{}).Count(&totalUsers)

	// Revenue (only successful orders)
	database.DB.Model(&order_entity.Order{}).
		Where("status = ?", "delivered").
		Select("COALESCE(SUM(total_amount),0)").
		Scan(&totalRevenue)

	c.JSON(200, gin.H{
		"totalOrders":      totalOrders,
		"pendingOrders":    pendingOrders,
		"successfulOrders": successfulOrders,
		"totalUsers":       totalUsers,
		"totalRevenue":     totalRevenue,
	})
}
