package middleware

import (
	"backend/internal/auth/entity"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckBlockedUser(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userID := ctx.GetUint("user_id")

		var user entity.User

		err := db.Select("is_blocked").First(&user, userID).Error

		if err != nil || user.IsBlocked {
			ctx.AbortWithStatusJSON(403, gin.H{"error": "User Is Blocked"})
			return

		}
 ctx.Next()
	}
}
