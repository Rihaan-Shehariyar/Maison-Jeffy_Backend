package address_routes

import (
	address_handler "backend/internal/address/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AddressRoutes(r *gin.Engine, h *address_handler.AddressHandler) {

	addr := r.Group("/address")
	addr.Use(middleware.JWTAuth())

	addr.POST("", h.Create)
	addr.GET("", h.GetMyAddress)
	addr.PUT("/:id", h.Update)
	addr.DELETE("/:id", h.Delete)

}
