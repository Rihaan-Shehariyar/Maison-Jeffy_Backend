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
  
 token,err:=jwt.ParseWithClaims(tokenstr,jwt.MapClaims{},func(t *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("JWT_ACCESS_TOKEN")),nil
})

 if err!=nil || !token.Valid {
	response.Unauthorized(ctx,"Invalid or Token Expired")
    ctx.Abort()
     return 
 }

 claims:=token.Claims.(*jwtutils.Claims)
 ctx.Set("user_id",claims.UserID)
 ctx.Set("email",claims.Email)


 ctx.Next()

}
}