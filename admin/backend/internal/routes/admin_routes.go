package routes

import (
	"backend/admin/backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {

	admin := r.Group("/admin")

    // Login Api
   
	admin.POST("/login", handlers.AdminLogin)
    
   // Middleware Authentication
   
	admin.Use(middleware.JWTAuth(), middleware.AdminOnly())

    // Product api

	admin.POST("/products", handlers.CreateProduct)
	admin.GET("/products", handlers.GetAllProducts)
	admin.DELETE("/products/:id", handlers.DeleteProduct)
	admin.PUT("/products/:id", handlers.UpdateProduct)

  // User Api
    
	admin.GET("/users", handlers.GetAllUsers)
	admin.PUT("/users/:id/block", handlers.BlockUser)

  // Order api
	admin.GET("/orders", handlers.GetAllOrders)
	admin.PUT("/orders/:id/status", handlers.UpdateOrderStatus)

   // Dashboard api
	admin.GET("/dashboard", handlers.GetDashboardStats)

}
