package profile_usecase

import (
	"backend/internal/auth/entity"
	"backend/internal/auth/repository"
	"errors"
)

type ProfileUseCase struct {
	userRepo repository.UserRepository
}

func NewProfileUseCase (userRepo repository.UserRepository)*ProfileUseCase{
    return &ProfileUseCase{ userRepo}
}


func(u *ProfileUseCase) GetProfile(userId uint)(*entity.User,error){

  user,err:= u.userRepo.FindByID(userId)
    if err!=nil || user == nil{
    return nil,errors.New("User Not Found")
}

 return user,nil
}

 func (u *ProfileUseCase) UpdateName(userId uint,name string)error{

   if name ==""{
   return  errors.New("Name Cannot Be empty")  
}

 return u.userRepo.UpdateName(userId,name)

 } 

