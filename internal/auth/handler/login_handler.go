package handler

import (
	"backend/internal/auth/usecase"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	usecase *usecase.LoginUsecase
}

func NewLoginHandler(u *usecase.LoginUsecase) *LoginHandler {
	return &LoginHandler{usecase: u}
}

func (h *LoginHandler) Login(c *gin.Context) {

	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid request")
		return
	}
	access, refresh, err := h.usecase.Login(req.Email, req.Password)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message":       "Login Successfull",
		"access_token":  access,
		"refresh_token": refresh,
	})

}
