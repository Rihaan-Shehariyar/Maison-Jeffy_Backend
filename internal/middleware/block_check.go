package middleware

import (
	"backend/internal/auth/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckBlockedUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		userID, exists := ctx.Get("user_id")
		if !exists {
			ctx.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}

		var user entity.User

		err := db.Select("is_blocked").
			First(&user, userID.(uint)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(404, gin.H{"error": "User not found"})
			return
		}

		if user.IsBlocked {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "User Is Blocked"})
			return
		}

		ctx.Next()
	}
}
