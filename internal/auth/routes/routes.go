package routes

import (
	"backend/internal/auth/handler"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
r *gin.RouterGroup,
h *handler.AuthHandler,
l *handler.LoginHandler,
refresh *handler.RefreshHandler,
forgot *handler.ForgetPasswordHandler,
reset *handler.ResetPasswordHandler, 
 ){
  r.POST("/signup",h.Signup)
  r.POST("/login",l.Login)
  r.POST("/refresh",refresh.Refresh)
  r.POST("/forgot-password",forgot.ForgetPassword)
  r.POST("/reset-password",reset.ResetPassword)
}

func OTPRoutes(r *gin.RouterGroup,h *handler.OTPHandler){
  r.POST("/send-otp",h.SendOTP)
  r.POST("/verify-otp",h.VerifyOTP)
}