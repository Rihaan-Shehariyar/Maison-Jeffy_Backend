package repository

import "backend/internal/auth/entity"

type UserRepository interface {
	FindByEmail(email string) (*entity.User, error)
	Create(user *entity.User) error
}