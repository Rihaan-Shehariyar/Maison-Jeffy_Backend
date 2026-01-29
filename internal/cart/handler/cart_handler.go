package cart_handler

import (
	cart_usecase "backend/internal/cart/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CartHandler struct {
	usecase *cart_usecase.CartUsecase
}

func NewCartHandler(usecase *cart_usecase.CartUsecase) *CartHandler {
	return &CartHandler{usecase}
}

func (h *CartHandler) Add(c *gin.Context) {

	userID := c.GetUint("user_id")

	idParam := c.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid Product ID")
		return
	}

	if err := h.usecase.Add(userID, uint(productID)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}
	c.JSON(200, gin.H{"message": "Product Added To Cart"})
}

func (h *CartHandler) Remove(c *gin.Context) {

	userID := c.GetUint("user_id")

	idParam := c.Param("productID")
	productID, err := strconv.Atoi(idParam)
	if err != nil {
		response.BadRequest(c, "Invalid Product ID")
		return
	}

	if err := h.usecase.Remove(userID, uint(productID)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Product removed From Cart"})

}

func (h *CartHandler) UpdateQty(c *gin.Context) {

	type updarequantity struct {
		Quantity int `json:"quantity"`
	}

	userID := c.GetUint("user_id")
	idParam := c.Param("productID")

	productID, err := strconv.Atoi(idParam)

	if err != nil {
		response.BadRequest(c, "Invalid Product ID")
		return
	}

	var req updarequantity

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Ivavlid JSON")
		return
	}

	if req.Quantity <= 0 {
		response.BadRequest(c, "Quantity Cannot be negative")
		return
	}

	if err := h.usecase.UpdateQty(userID, uint(productID), req.Quantity); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Quantity Updated"})

}

func (h *CartHandler) GetMyCart(c *gin.Context) {

	userID := c.GetUint("user_id")

	items, err := h.usecase.GetMyUser(userID)
	if err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, items)

}

func (h *CartHandler) Clear(c *gin.Context) {

	userID := c.GetUint("user_id")

	if err := h.usecase.Clear(userID); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Products cleared from Cart"})

}
