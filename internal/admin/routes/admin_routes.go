package admin_routes

import (
	admin_handler "backend/internal/admin/handler"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine,h *admin_handler.UserAdminHandler){

 admin := r.Group("/admin")
 r.Use(middleware.JWTAuth(),middleware.AdminOnly())

 users := admin.Group("/users")
 users.PUT("/:id",h.UpdateUser)
 users.PUT("/:id",h.BlockUser)
}