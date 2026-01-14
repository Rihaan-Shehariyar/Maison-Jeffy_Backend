package repository

import (
	"backend/internal/auth/entity"
	"errors"

	"gorm.io/gorm"
)

type userRepositoryPg struct {
	db *gorm.DB
}

type userModel struct{
 
  Id uint `gorm:"primaryKey"`
  Name string
  Email string `gorm:"uniqueIndex"`
  Password string
  Role string
  IsActive bool

}


func(r *userRepositoryPg)FindByEmail(email string)(*entity.User,error){

  var u userModel

 err:=r.db.Where("email=?",u.Email).Error

 if errors.Is(err,gorm.ErrRecordNotFound){
  return nil,nil
}


 if err!=nil{
   return nil,nil
}

 return &entity.User{
  ID: u.Id,
  Name: u.Name,
  Email: u.Email,
  Password: u.Password,
  Role: u.Role,
  IsActive: u.IsActive,
},nil


}


func (r *userRepositoryPg)Create(user *entity.User)(error){

  u:=userModel{
 Name: user.Name,
 Email: user.Email,
 Password: user.Password,
 Role: user.Role,
 IsActive: user.IsActive,

}

 return r.db.Create(&u).Error

}