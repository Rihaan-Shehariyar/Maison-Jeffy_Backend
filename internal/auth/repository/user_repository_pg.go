package repository

import (
	"backend/internal/auth/entity"
	"errors"

	"gorm.io/gorm"
)

type userRepositoryPg struct {
	db *gorm.DB
}

func NewUserRepositoryPg(db *gorm.DB)UserRepository{
   return &userRepositoryPg{db: db}
}

 




func(r *userRepositoryPg)FindByEmail(email string)(*entity.User,error){

  var u entity.User

 err:=r.db.Where("email = ?",email).First(&u).Error

 if errors.Is(err,gorm.ErrRecordNotFound){
  return nil,nil
}


 if err!=nil{
   return nil,nil
}

 return &entity.User{
  ID: u.ID,
  Name: u.Name,
  Email: u.Email,
  Password: u.Password,
  Role: u.Role,
  IsVerified: u.IsVerified,
},nil


}


func (r *userRepositoryPg)Create(user *entity.User)(error){

  u:=entity.User{
 
 Name: user.Name,
 Email: user.Email,
 Password: user.Password,
 Role: user.Role,
 IsVerified: user.IsVerified,

}

 return r.db.Create(&u).Error

}

func (r *userRepositoryPg)MarkVerfied(email string)error{
   return r.db.Model(&entity.User{}).
           Where("emai = ?",email).
          Update("isVerified",true).Error
}