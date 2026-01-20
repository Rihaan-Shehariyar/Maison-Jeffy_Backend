package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Created(c *gin.Context,msg string,data interface{}){
 c.JSON(200,gin.H{
   "status" : "success",
   "message" : msg,
    "data" : data,
})
}

func BadRequest(c *gin.Context,msg string){
  c.JSON(400,gin.H{
   "status" : "error",
    "message" : msg,
})
}

func Conflict(c *gin.Context,msg string){
  c.JSON(409,gin.H{
  "status":"error",
   "error":msg, 
})
}

func InternalError(c *gin.Context,msg string){
  c.JSON(500,gin.H{"error":"Internal Server Error",
                 "status" : msg, 
})
}

func Unauthorized(c *gin.Context,msg string){
   c.JSON(http.StatusUnauthorized,gin.H{
   "status" : "error",
    "error" : msg,
})
}