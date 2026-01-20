package main

import (
	admin_handler "backend/internal/admin/handler"
	admin_routes "backend/internal/admin/routes"
	admin_usecase "backend/internal/admin/usecase"
	"backend/internal/auth/entity"
	"backend/internal/auth/handler"
	"backend/internal/auth/repository"
	"backend/internal/auth/routes"
	"backend/internal/auth/usecase"
	"backend/internal/middleware"
	entitys "backend/internal/product/entity"
	"backend/internal/product/handlers"
	"backend/internal/product/repositorys"
	product_routes "backend/internal/product/routes"
	"backend/internal/product/usecase"
	profile_handler "backend/internal/user/handler"
	profile_routes "backend/internal/user/routes"
	profile_usecase "backend/internal/user/usecase"
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
&entitys.Product{},
);err!=nil{
  log.Fatal("AutoMigrate failed:", err)
 }

 
// Repository
 user_repo := repository.NewUserRepositoryPg(db)
 otp_repo := repository.NewOTPRepository(db)

// USE_CASE
 signup_uc := usecase.NewSignupCase(user_repo)
 otp_uc := usecase.NewOTPUsecase(otp_repo,user_repo)
 login_uc:=usecase.NewLoginUseCase(user_repo)
 refresh_uc :=usecase.NewRefreshUseCase()
 forgot_uc :=usecase.NewForgetPasswordUseCase(user_repo,otp_repo)
 reset_uc := usecase.NewResetPasswordUsecase(user_repo,otp_repo)
 
// HANDLERS
 auth_handler:= handler.NewAuthHandler(signup_uc)
 otp_handler := handler.NewOTPHandler(otp_uc)
 login_handler := handler.NewLoginHandler(login_uc)
 refresh_handler := handler.NewRefreshHandler(refresh_uc)
 forgot_handler := handler.NewForgetPasswordHandler(forgot_uc)
 reset_handler := handler.NewResetPasswordHandler(reset_uc)



// Products

productRepo := repositorys.NewProductRepositoryPg(db)
product_uc := usecases.NewProductRepositoryUseCase(productRepo)
product_handler:=handlers.NewProductHandler(product_uc)

// User_Profiles

profile_uc := profile_usecase.NewProfileUseCase(user_repo)
profile_handler := profile_handler.NewProfileGHandler(profile_uc)

//admin

admin_uc := admin_usecase.NewUserAdminUscase(user_repo)
adminHandler := admin_handler.NewUserAdminHandler(admin_uc)

admin_product_uc := admin_usecase.NewProductAdminUsecase(productRepo)
admin_product_handler := admin_handler.NewProductAdminHandler(admin_product_uc)


// Router

 r:=gin.Default()

 api:=r.Group("/auth")
 routes.RegisterRoutes(api,auth_handler,login_handler,refresh_handler,forgot_handler,reset_handler)
 routes.OTPRoutes(api,otp_handler)


protected := api.Group("")
protected.Use(middleware.JWTAuth())

 
 protected.GET("/profile",func(ctx *gin.Context) {
   ctx.JSON(200,gin.H{
    "user_id" : ctx.GetUint("user_id"),
     "email" : ctx.GetString("email"),
})
 })

  // Product_routes

 product_routes.ProductRoutes(r,product_handler)

// Profile_routes

profile_routes.UserRoutes(r,profile_handler)

// admin_routes

admin_routes.AdminRoutes(r,adminHandler,admin_product_handler)

 log.Println("server running on :8080")
 if err:=r.Run(":8080");err!=nil{
  log.Fatal(err)
}

}