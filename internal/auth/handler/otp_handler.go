package handler

import (
	"backend/internal/auth/usecase"
	"backend/pkg/response"
	"github.com/gin-gonic/gin"
)

type OTPHandler struct {
	usecase *usecase.OTPUsecase
}

func NewOTPHandler(u *usecase.OTPUsecase)*OTPHandler{

 return &OTPHandler{
   usecase: u,
}

}

// Send-Otp

func (h *OTPHandler)SendOTP(c *gin.Context){
  
 var req struct{
   Email string `json:"email"`
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
   response.BadRequest(c,"Invalid request")
   return
} 

 if err:= h.usecase.SendOTP(req.Email);err!=nil{
  response.BadRequest(c,err.Error())
   return
}

c.JSON(200,gin.H{"message":"OTP Sent"})
 

}


// Verify-Otp

func (h *OTPHandler) VerifyOTP(c *gin.Context){

 var req struct{
  
  Email string `json:"email"`
  OTP string `json:"otp"`

}   

 if err:=c.ShouldBindJSON(&req);err!=nil{
   response.BadRequest(c,"Invalid request")
   return
}

 if err:=h.usecase.VerifyOTP(req.Email,req.OTP);err!=nil{
   response.BadRequest(c,"Invalid OTP")
   return
} 

 c.JSON(200,gin.H{"message":"Email verified Successfully"})


}