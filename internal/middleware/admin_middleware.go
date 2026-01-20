package middleware

import "github.com/gin-gonic/gin"

func AdminOnly() gin.HandlerFunc{
 return func(ctx *gin.Context) {
   role := ctx.GetString("role")
 if role!= "admin"{
  ctx.JSON(403,gin.H{"error":"Admin Access required "})
  ctx.Abort() 
   return 
}  

ctx.Next()
}
}