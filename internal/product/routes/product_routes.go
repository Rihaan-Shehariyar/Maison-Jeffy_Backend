package product_routes

import (
	"backend/internal/middleware"
	"backend/internal/product/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, h *handlers.ProductHandler) {
	r.GET("/products", h.GetAllProducts)
	r.GET("/products/:id", h.GetProductById)

	admin := r.Group("/products")

	admin.Use(middleware.JWTAuth(), middleware.AdminOnly())

}
