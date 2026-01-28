package wishlist_routes

import (
	"backend/internal/middleware"
	wishlist_handler "backend/internal/wishlist/handler"

	"github.com/gin-gonic/gin"
)

func WishlistRoutes(r *gin.Engine,h *wishlist_handler.WishlistHandler){

 w := r.Group("/wishlist")
 w.Use(middleware.JWTAuth())

 w.POST("/:productID",h.Add)
 w.DELETE("/:productID",h.Remove)
 w.GET("",h.GetMyWishlist)

}