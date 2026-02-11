package main

import (
	"log"
	"time"

	"backend/admin/backend/internal/database"
	"backend/admin/backend/internal/models"
	"backend/admin/backend/internal/routes"
	entitys "backend/internal/product/entity"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// DB
	database.Connect()
	database.DB.AutoMigrate(
		&models.Admin{},
		&entitys.Product{},
	)
	database.SeedAdmin()

	// Gin engine
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	//  CORS MUST BE HERE 
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://127.0.0.1:5500",
			"http://localhost:5500",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type", "Authorization",
		},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	// Routes AFTER CORS
	routes.AdminRoutes(r)

	// Static files
	r.Static("/uploads", "./uploads")


	log.Println("ADMIN SERVER RUNNING ON :8080")
	r.Run(":8080")
}
