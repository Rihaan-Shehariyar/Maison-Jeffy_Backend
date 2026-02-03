package handler

import (
	"backend/internal/auth/usecase"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)

type ResetPasswordHandler struct {
	usecase *usecase.ResetPasswordUsecase
}

func NewResetPasswordHandler(u *usecase.ResetPasswordUsecase) *ResetPasswordHandler{
 return &ResetPasswordHandler{usecase: u }
}


func (h *ResetPasswordHandler) ResetPassword(c *gin.Context){

  var req struct{
   Email string `json:"email"`
   OTP string `json:"otp"`
   NewPassword string `json:"new_password"`
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
  response.BadRequest(c,"Invalid Json")
 return
}

 if err:=h.usecase.ResetPassword(req.Email,req.OTP,req.NewPassword);err!=nil{
    response.BadRequest(c,err.Error())
   return
}

 c.JSON(200,gin.H{"message":"Password Reset Successfully"})

}



