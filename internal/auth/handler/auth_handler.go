package handler

import (
	"backend/internal/auth/dto"
	"backend/internal/auth/usecase"
	"backend/pkg/response"

	"github.com/gin-gonic/gin"
)



type AuthHandler struct {
	signupUC *usecase.SignupUseCase
}

func NewAuthHandler(signupUc *usecase.SignupUseCase) *AuthHandler{
  
  return &AuthHandler{signupUC: signupUc}
}

func (h *AuthHandler)Signup(c *gin.Context){

  var req dto.SignupRequest

 if err:= c.ShouldBindJSON(&req);err!=nil{
  response.BadRequest(c,err.Error())
  return
}


 err:=h.signupUC.Signup(req.Name,req.Email,req.Password)

 if err!=nil{
   if err ==usecase.ErrEmailAlreadyExists{
  response.Conflict(c,"email already exist")
  return
}

 response.InternalError(c,"Server Error")
 return
}

response.Created(c,"Signup Succesfully,verify otp",nil)

}