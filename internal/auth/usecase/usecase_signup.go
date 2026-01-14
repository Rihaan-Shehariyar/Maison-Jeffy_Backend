package usecase

import (
	"backend/internal/auth/entity"
	"backend/internal/auth/repository"
	"errors"


	"golang.org/x/crypto/bcrypt"
)

var ErrEmailAlreadyExists = errors.New("email already registered")

type SignupUseCase struct {
	userRepo repository.UserRepository
}

func NewSignupCase(userRepo repository.UserRepository)*SignupUseCase{
   return &SignupUseCase{userRepo: userRepo}
}


func (u *SignupUseCase)Signup(name,email,password string)error{
 
  existing,err:=u.userRepo.FindByEmail(email)
 if err!=nil{
  return err
}

 if existing!=nil{
  return ErrEmailAlreadyExists
}
 
 hashed,err := bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
 if err!=nil{
   return err
}

 user:=entity.User{
 Name: name,
 Email: email,
 Password: string(hashed),
 Role: "user",
 IsActive: false,
}

return u.userRepo.Create(&user)

}

