package wishlist_handler

import (
	wishlist_usecase "backend/internal/wishlist/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type WishlistHandler struct {
	usecase *wishlist_usecase.WishlistUsecase
}

func NewWishlistHandler(usecase *wishlist_usecase.WishlistUsecase) *WishlistHandler {
	return &WishlistHandler{usecase}
}

func (h *WishlistHandler) Add(c *gin.Context) {

	userID := c.GetUint("user_id")
	idParam := c.Param("productID")

	if idParam == "" {

		response.BadRequest(c, "product is missing")
		return
	}

	productID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil || productID == 0 {
		response.BadRequest(c, err.Error())
		return
	}

	if err := h.usecase.Add(userID, uint(productID)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Added To WishList"})

}

func (h *WishlistHandler) Remove(c *gin.Context) {

	userID := c.GetUint("user_id")
	pID, _ := strconv.Atoi(c.Param("productID"))

	h.usecase.Remove(userID, uint(pID))

	c.JSON(200, gin.H{"message": "Product Remove From Wishlist"})

}

func (h *WishlistHandler) GetMyWishlist(c *gin.Context) {

	userID := c.GetUint("user_id")
	items, _ := h.usecase.GetMyWishlist(userID)
	c.JSON(200, items)

}
