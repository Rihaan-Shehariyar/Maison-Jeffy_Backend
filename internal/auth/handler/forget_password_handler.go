package handler

import (
	"backend/internal/auth/usecase"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ForgetPasswordHandler struct {
	usecase *usecase.ForgetPasswordUsecase
}

func NewForgetPasswordHandler(u *usecase.ForgetPasswordUsecase) *ForgetPasswordHandler {
	return &ForgetPasswordHandler{usecase: u}
}


// Forgot-Password
func (h *ForgetPasswordHandler) ForgetPassword(c *gin.Context) {

	var req struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid Json")
		return
	}

	if err := h.usecase.SendResendOTP(req.Email); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "OTP has Sent to renew Password"})

}
