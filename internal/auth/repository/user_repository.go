package repository

import (
	"backend/internal/auth/entity"

)

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
    FindByID(id uint)(*entity.User,error)
	Create(user *entity.User) error
    MarkVerfied(email string) error
    UpdatePassword(email,password string) error
    UpdateName(id uint,name string)error
    UpdateUser(user *entity.User)error
    BlockUser(id uint)error

}

