package database

import (
	"backend/internal/auth/entity"
     
)

func SeedAdmin() {

	var count int64

	DB.Model(&entity.User{}).Count(&count)

	if count == 0 {
		DB.Create(&entity.User{
			Email:    "admin@gmail.com",
			Password: "Admin@123",
            Role: "admin",
            
		})
	}

}
