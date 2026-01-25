package admin_usecase

import (
	"backend/internal/auth/repository"
	"errors"
)

type UserAdminUsecase struct {
	userRepo repository.UserRepository
}

func NewUserAdminUscase(userRepo repository.UserRepository) *UserAdminUsecase {
	return &UserAdminUsecase{userRepo}
}

func (u *UserAdminUsecase) UpdateUser(id uint, name string, email string, role string) error {

	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return errors.New("User Not Found")
	}

	if user.Role == "admin" {
		return errors.New("Cannot modify user admin")
	}

	if name != "" {
		user.Name = name
	}

	if email != "" {

		if email != user.Email {
			existing, _ := u.userRepo.FindByEmail(email)
			if existing != nil && existing.ID != user.ID {

				return errors.New("email already in exist")
			}

			user.Email = email
		}

	}

	if role != "" {
		if role != "user" && role != "admin" {
			return errors.New("Invalid Role")
		}
		user.Role = role
	}

	return u.userRepo.UpdateUser(user)
}

func (u *UserAdminUsecase) BlockUser(id uint) error {

	user, err := u.userRepo.FindByID(id)
	if err != nil || user == nil {
		return errors.New("User Not Found")
	}

	if user.Role == "admin" {
		return errors.New("Cannot Block Admin")
	}

	return u.userRepo.BlockUser(id)

}
