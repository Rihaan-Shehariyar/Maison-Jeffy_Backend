package handler

import (
	"backend/internal/auth/usecase"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type RefreshHandler struct {
	usecase *usecase.RefreshUseCase
}
func NewRefreshHandler(u *usecase.RefreshUseCase)*RefreshHandler{
  return &RefreshHandler{usecase: u}
}

func(h *RefreshHandler)Refresh(c *gin.Context){
   var req struct{
    RefreshToken string `json:"refresh_token"`
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
    response.BadRequest(c,"Invalid Json Format")
     return
}

 access_token,err:=h.usecase.Refresh(req.RefreshToken) 
 if err!=nil{
    response.Unauthorized(c,err.Error())
}

 c.JSON(200,gin.H{"access_token" : access_token})

}

