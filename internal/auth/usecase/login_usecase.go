package usecase

import (
	"backend/internal/auth/repository"
	jwtutils "backend/pkg/jwt_utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userRepo repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository)*LoginUsecase{
    return &LoginUsecase{userRepo: userRepo}
}

func (u *LoginUsecase) Login(email,password string)(string,string,error){

 

  user,err:= u.userRepo.FindByEmail(email)
   if err!=nil{
   return "","",err
}

// If user is nill
if user==nil{
  return "","",errors.New("Invalid email")
}

 if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));
err!=nil{
   return "","",errors.New("Invalid Password")
}

 if !user.IsVerified{
   return "","",errors.New("Please Verify Your email Before Login")
}

 access_token,err:=jwtutils.GenerateAccessToken(user.ID,user.Email)
 if err!=nil{
    return "","",err
}

refresh_token,err := jwtutils.GenerateRefreshToken(user.Email,user.ID)
 if err!=nil{
    return "","",err
}

 return access_token,refresh_token,nil

}

