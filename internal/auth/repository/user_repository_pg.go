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


func(r* userRepositoryPg)FindByID(userId uint)(*entity.User,error){

  var u entity.User

 if err:=r.db.First(&u,userId).Error;err!=nil{
    return nil,err
}

 return &entity.User{
   ID: u.ID,
   Name: u.Name,
   Email: u.Email,
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
           Where("email = ?",email).
          Update("isVerified",true).Error
}


func (r userRepositoryPg)UpdatePassword(email,password string)error{
  
 return r.db.Model(&entity.User{}).
         Where("email = ? ",email).
         Update("password",password).Error

}


func (r *userRepositoryPg)UpdateName(id uint,name string)error{

 return r.db.Model(&entity.User{}).
             Where("id = ?",id).
             Update("name",name).Error
  
}

func (r *userRepositoryPg)UpdateUser(user *entity.User)error{
   return r.db.Save(user).Error
}


func (r *userRepositoryPg)BlockUser(id uint)error{
 return r.db.Model(&entity.User{}).
             Where("id = ?",id).
             Update("IsBlocked",true).Error 
}

