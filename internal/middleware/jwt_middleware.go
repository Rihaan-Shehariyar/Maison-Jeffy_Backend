package middleware

import (
	jwtutils "backend/pkg/jwt_utils"
	"backend/pkg/response"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc{
 return func(ctx *gin.Context) {

   auth:=ctx.GetHeader("Authorization")
  if auth =="" {
	response.Unauthorized(ctx,"Token Missing")
    ctx.Abort()
     return 
  }

 tokenstr:=strings.TrimPrefix(auth, "Bearer ")
 claims := &jwtutils.Claims{}
  
 token,err:=jwt.ParseWithClaims(tokenstr,claims,func(t *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("JWT_SECRET")),nil
})

 if err!=nil || !token.Valid {
	response.Unauthorized(ctx,"Invalid or Token Expired")
    ctx.Abort()
     return 
 }

 ctx.Set("user_id",claims.UserID)
 ctx.Set("email",claims.Email)
 ctx.Set("role",claims.Role)

 ctx.Next()

}
}