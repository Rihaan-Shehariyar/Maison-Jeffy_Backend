package handlers

import (
	"backend/admin/backend/internal/database"
	"backend/internal/auth/entity"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Get Users

func GetAllUsers(c *gin.Context) {

	var users []entity.User

	if err := database.DB.Find(&users).Error; err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(200, users)

}

// Block Users

func BlockUser(c *gin.Context) {

	id := c.Param("id")

	var user entity.User

	if err := database.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User Not Found"})
		return
	}

	user.IsBlocked = !user.IsBlocked

	if err := database.DB.Save(&user).Error; err != nil {
		response.InternalError(c, "Failed to update User")
		return
	}

	c.JSON(200, gin.H{
		"id":         user.ID,
		"is_blocked": user.IsBlocked,
	})

}
