package routes

import (
	"backend/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup,h *handler.AuthHandler){
  r.POST("signup",h.Signup)
}

func OTPRoutes(r *gin.RouterGroup,h *handler.OTPHandler){
  r.POST("send-otp",h.SendOTP)
  r.POST("verify-otp",h.VerifyOTP)
}