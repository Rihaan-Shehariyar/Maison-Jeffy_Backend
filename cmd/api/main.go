package main

import (
	"backend/internal/auth/entity"
	"backend/internal/auth/handler"
	"backend/internal/auth/repository"
	"backend/internal/auth/routes"
	"backend/internal/auth/usecase"
	"backend/pkg/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err:= godotenv.Load();err!=nil{
    log.Printf("no .env file found")
 }

 dsn:= os.Getenv("POSTGRES_DSN")
 db,err := database.Connect(dsn)
 if err!=nil{
    log.Fatal("Failed to Connect DB : ",err)
}


if err:=db.AutoMigrate(
&entity.User{},
&entity.OTP{},
);err!=nil{
  log.Fatal("AutoMigrate failed:", err)
 }

 
// Repository
 user_repo := repository.NewUserRepositoryPg(db)
 otp_repo := repository.NewOTPRepository(db)

// USE_CASE
 signup_uc := usecase.NewSignupCase(user_repo)
 otp_uc := usecase.NewOTPUsecase(otp_repo,user_repo)
 
// HANDLERS
 auth_handler:= handler.NewAuthHandler(signup_uc)
 otp_handler := handler.NewOTPHandler(otp_uc)

 r:=gin.Default()

 api:=r.Group("/auth")
 routes.RegisterRoutes(api,auth_handler)
 routes.OTPRoutes(api,otp_handler)

 log.Println("server running on :8080")
 if err:=r.Run(":8080");err!=nil{
  log.Fatal(err)
}

}