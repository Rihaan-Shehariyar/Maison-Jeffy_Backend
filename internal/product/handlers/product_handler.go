package handlers

import (
	usecases "backend/internal/product/usecase"
	"backend/pkg/response"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	usecase *usecases.ProductUseCase
}

func NewProductHandler(u *usecases.ProductUseCase) *ProductHandler {
	return &ProductHandler{u}
}

// Get All Products

func (h *ProductHandler) GetAllProducts(c *gin.Context) {

	products, err := h.usecase.GetAllProducts()
	if err != nil {
		response.InternalError(c, "Failed To Fetch Products")
		return
	}

	c.JSON(200, products)
}

// Get Products By Id

func (h *ProductHandler) GetProductById(c *gin.Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	products, err := h.usecase.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product Not Found"})
		return
	}

	c.JSON(200, products)

}
