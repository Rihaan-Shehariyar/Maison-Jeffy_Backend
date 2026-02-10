package handlers

import (
	"backend/admin/backend/internal/database"
	entitys "backend/internal/product/entity"
	"backend/pkg/response"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context) {

	var products []entitys.Product
	database.DB.Find(&products)
	c.JSON(200, products)

}

func CreateProduct(c *gin.Context) {

	name := c.PostForm("name")
	priceStr := c.PostForm("price")
	stockstr := c.PostForm("stock")
	description := c.PostForm("description")
	category := c.PostForm("category")
	sku := c.PostForm("sku")

	if name == "" || priceStr == "" || stockstr == "" || sku == "" {
		response.BadRequest(c, "Name,price,SKU and Stock is required")
		return
	}

	var count int64
	database.DB.Model(&entitys.Product{}).
		Where("sku = ?", sku).
		Count(&count)

	if count > 0 {
		response.BadRequest(c, "SKU already exists")
		return
	}

	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil || price < 0 {
		response.BadRequest(c, "Invalid Price")
		return
	}

	stock, err := strconv.Atoi(stockstr)
	if err != nil || stock < 0 {
		response.BadRequest(c, "Invalid Price")
		return
	}

	file, err := c.FormFile("image")
	if err != nil {
		response.BadRequest(c, "Image is required")
		return
	}

	uploadDir := "uploads/products"
	_ = os.Mkdir(uploadDir, os.ModePerm)

	fileName := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
	imagePath := filepath.Join(uploadDir, fileName)

	if err := c.SaveUploadedFile(file, imagePath); err != nil {
		response.InternalError(c, "Failed to Save Image")
		return
	}

	product := entitys.Product{
		Name:        name,
		Description: description,
		Price:       price,
		Stock:       stock,
		ImageURL:    imagePath,
		Category:    category,
		SKU:         sku,
	}

	if err := database.DB.Create(&product).Error; err != nil {
		response.InternalError(c, "Failed To create Product")
		return
	}

	c.JSON(200, gin.H{
		"message": "Product Created Succeesfullt",
		"product": product,
	})

}

func UpdateProduct(c *gin.Context) {

	id := c.Param("id")

	var product entitys.Product

	if err := database.DB.First(&product, id).Error; err != nil {
		response.BadRequest(c, "not found")
		return
	}

	product.Name = c.PostForm("name")
	product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	product.Stock, _ = strconv.Atoi(c.PostForm("stock"))
	product.Category = c.PostForm("category")
	product.Description = c.PostForm("description")
	product.SKU = c.PostForm("sku")
	database.DB.Save(&product)
	c.JSON(200, product)

}

func DeleteProduct(c *gin.Context) {

	database.DB.Delete(&entitys.Product{}, c.Param("id"))
	c.JSON(200, gin.H{"message": "Product Deleted Successfully"})

}
