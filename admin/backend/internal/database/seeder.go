package database

import "backend/admin/backend/internal/models"

func seedAdmin() {

	var count int64

	DB.Model(&models.Admin{}).Count(&count)

	if count == 0 {
		DB.Create(&models.Admin{
			Email:    "admin@gmail.com",
			Password: "Admin@123",
		})
	}

}
