package handlers

import (
	"backend/admin/backend/internal/database"
	"backend/internal/auth/entity"
	jwtutils "backend/pkg/jwt_utils"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid Json")
		return

	}

	var admin entity.User

	err := database.DB.Where("email = ? AND password = ? And role = admin", req.Email, req.Password).First(&admin).Error

	if err != nil {
		response.Unauthorized(c, "Invalid Credentials")
		return
	}

	token, err := jwtutils.GenerateAccessToken(admin.ID, admin.Email, "admin")
	if err != nil {
		response.InternalError(c, "Token Generation Failed")
		return
	}

	c.JSON(200, gin.H{"token": token})

}
