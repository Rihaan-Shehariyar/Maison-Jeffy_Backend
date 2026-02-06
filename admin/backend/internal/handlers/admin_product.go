package handlers

import (
	"backend/admin/backend/internal/database"
	entitys "backend/internal/product/entity"

	"github.com/gin-gonic/gin"
)

func GetAllProducts(c *gin.Context){
 
 var products []entitys.Product
 database.DB.Find(&products)
 c.JSON(200,products)

}

