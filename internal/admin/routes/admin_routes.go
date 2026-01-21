package admin_routes

import (
	admin_handler "backend/internal/admin/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine,h *admin_handler.UserAdminHandler,p *admin_handler.ProductAdminHandler){

 admin := r.Group("/admin")
 admin.Use(middleware.JWTAuth(),middleware.AdminOnly())

 users := admin.Group("/users")
 users.PUT("/:id",h.UpdateUser)
 users.PUT("/:id/block",h.BlockUser)

 products:=admin.Group("/products")
 products.POST("",p.CreateProduct)
 products.PUT("/:id",p.UpdateProduct)
 products.DELETE("/:id",p.DeleteProduct)
 
}