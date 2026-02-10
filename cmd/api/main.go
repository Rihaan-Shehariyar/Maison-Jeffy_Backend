package main

import (
	address_entity "backend/internal/address/entity"
	address_handler "backend/internal/address/handler"
	address_repository "backend/internal/address/repository"
	address_routes "backend/internal/address/routes"
	address_usecase "backend/internal/address/usecase"
	admin_handler "backend/internal/admin/handler"
	admin_routes "backend/internal/admin/routes"
	admin_usecase "backend/internal/admin/usecase"
	"backend/internal/auth/entity"
	"backend/internal/auth/handler"
	"backend/internal/auth/repository"
	"backend/internal/auth/routes"
	"backend/internal/auth/usecase"
	cart_entity "backend/internal/cart/entity"
	cart_handler "backend/internal/cart/handler"
	cart_repository "backend/internal/cart/repository"
	cart_routes "backend/internal/cart/routes"
	cart_usecase "backend/internal/cart/usecase"
	"backend/internal/middleware"
	order_entity "backend/internal/orders/entity"
	order_handler "backend/internal/orders/handler"
	order_repository "backend/internal/orders/repository"
	order_routes "backend/internal/orders/routes"
	order_usecase "backend/internal/orders/usecase"
	entitys "backend/internal/product/entity"
	"backend/internal/product/handlers"
	"backend/internal/product/repositorys"
	product_routes "backend/internal/product/routes"
	usecases "backend/internal/product/usecase"
	profile_handler "backend/internal/user/handler"
	profile_routes "backend/internal/user/routes"
	profile_usecase "backend/internal/user/usecase"
	wishlist_entity "backend/internal/wishlist/entity"
	wishlist_handler "backend/internal/wishlist/handler"
	wishlist_repository "backend/internal/wishlist/repository"
	wishlist_routes "backend/internal/wishlist/routes"
	wishlist_usecase "backend/internal/wishlist/usecase"
	"backend/pkg/database"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("no .env file found")
	}

	dsn := os.Getenv("POSTGRES_DSN")
	db, err := database.Connect(dsn)
	if err != nil {
		log.Fatal("Failed to Connect DB : ", err)
	}

	if err := db.AutoMigrate(
		&entity.User{},
		&entity.OTP{},
		&entitys.Product{},
		&wishlist_entity.Wishlist{},
		&cart_entity.Cart{},
		&order_entity.Order{},
		&order_entity.OrderItem{},
		&address_entity.Address{},
	); err != nil {
		log.Fatal("AutoMigrate failed:", err)
	}

	// Auth

	// Repository
	user_repo := repository.NewUserRepositoryPg(db)
	otp_repo := repository.NewOTPRepository(db)

	// USE_CASE
	signup_uc := usecase.NewSignupCase(user_repo)
	otp_uc := usecase.NewOTPUsecase(otp_repo, user_repo)
	login_uc := usecase.NewLoginUseCase(user_repo)
	refresh_uc := usecase.NewRefreshUseCase()
	forgot_uc := usecase.NewForgetPasswordUseCase(user_repo, otp_repo)
	reset_uc := usecase.NewResetPasswordUsecase(user_repo, otp_repo)

	// HANDLERS
	auth_handler := handler.NewAuthHandler(signup_uc)
	otp_handler := handler.NewOTPHandler(otp_uc)
	login_handler := handler.NewLoginHandler(login_uc)
	refresh_handler := handler.NewRefreshHandler(refresh_uc)
	forgot_handler := handler.NewForgetPasswordHandler(forgot_uc)
	reset_handler := handler.NewResetPasswordHandler(reset_uc)

	// Products

	product_Repo := repositorys.NewProductRepositoryPg(db)
	product_uc := usecases.NewProductRepositoryUseCase(product_Repo)
	product_handler := handlers.NewProductHandler(product_uc)

	// User_Profiles

	profile_uc := profile_usecase.NewProfileUseCase(user_repo)
	profile_handler := profile_handler.NewProfileGHandler(profile_uc)

	//admin

	admin_uc := admin_usecase.NewUserAdminUscase(user_repo)
	adminHandler := admin_handler.NewUserAdminHandler(admin_uc)

	admin_product_uc := admin_usecase.NewProductAdminUsecase(product_Repo)
	admin_product_handler := admin_handler.NewProductAdminHandler(admin_product_uc)

	// Router

	r := gin.Default()

	api := r.Group("/auth")
    api.Use(middleware.CheckBlockedUser(db))
	routes.RegisterRoutes(api, auth_handler, login_handler, refresh_handler, forgot_handler, reset_handler)
	routes.OTPRoutes(api, otp_handler)

	protected := api.Group("")
	protected.Use(middleware.JWTAuth())

	protected.GET("/profile", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"user_id": ctx.GetUint("user_id"),
			"email":   ctx.GetString("email"),
		})
	})

	// Product_routes

	product_routes.ProductRoutes(r, product_handler)

	// Profile_routes

	profile_routes.UserRoutes(r, profile_handler)

	//wishlist

	wishlist_repo := wishlist_repository.NewWishlistRepositoryPg(db)
	wishlist_uc := wishlist_usecase.NewWishlistUsecasePg(wishlist_repo)
	wishlist_handler := wishlist_handler.NewWishlistHandler(wishlist_uc)

	wishlist_routes.WishlistRoutes(r, wishlist_handler)

	// Cart

	cart_repo := cart_repository.NewCartRepositoryPg(db)
	cart_uc := cart_usecase.NewCartUsecase(cart_repo)
	cart_handler := cart_handler.NewCartHandler(cart_uc)

	cart_routes.CartRoutes(r, cart_handler)

	// orders

	order_repo := order_repository.NewOrderRepositoryPg(db)
	order_uc := order_usecase.NewOrderUsecase(db, order_repo, cart_repo, product_Repo)
	order_handler := order_handler.NewOrderHandler(order_uc)

	order_routes.OrderRoutes(r, order_handler)

	// address

	address_repo := address_repository.NewAddressRepositoryPg(db)
	address_uc := address_usecase.NewAddressUsecase(address_repo)
	address_handler := address_handler.NewAddressHandler(address_uc)

	address_routes.AddressRoutes(r, address_handler)

	// admin_routes

	admin_routes.AdminRoutes(r, adminHandler, admin_product_handler, order_handler)

	// HTTP Server

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Println("Server running on : 8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Listen error :  %v\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("ShutDown Signal Recieved...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Forced To ShutDown", err)
	}

	sqlDB, _ := db.DB()
	sqlDB.Close()
	log.Println("Server Exited Cleanly")

}
