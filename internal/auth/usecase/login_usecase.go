package usecase

import (
	"backend/internal/auth/repository"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginUsecase struct {
	userRepo repository.UserRepository
}

func NewLoginUseCase(userRepo repository.UserRepository)*LoginUsecase{
    return &LoginUsecase{userRepo: userRepo}
}

func (u *LoginUsecase) Login(email,password string)error{

 

  user,err:= u.userRepo.FindByEmail(email)
   if err!=nil{
   return err
}

// If user is nill
if user==nil{
  return errors.New("Invalid email")
}

 if err:=bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password));
err!=nil{
   return errors.New("Invalid Password")
}

 if !user.IsVerified{
   return errors.New("Please Verify Your email Before Login")
}

return nil 
 
}

