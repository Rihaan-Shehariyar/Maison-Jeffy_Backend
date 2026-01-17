package routes

import (
	"backend/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup,h *handler.AuthHandler,l *handler.LoginHandler){
  r.POST("signup",h.Signup)
  r.POST("login",l.Login)
}

func OTPRoutes(r *gin.RouterGroup,h *handler.OTPHandler){
  r.POST("send-otp",h.SendOTP)
  r.POST("verify-otp",h.VerifyOTP)
}