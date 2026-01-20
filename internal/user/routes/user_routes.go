package profile_routes

import (
	"backend/internal/middleware"
	"backend/internal/user/handler"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine, h *profile_handler.ProfileHandler){

  user:=r.Group("/user")
  user.Use(middleware.JWTAuth())

 user.GET("/profile",h.GetProfile)
 user.PUT("/update_name",h.UpdateName)


}