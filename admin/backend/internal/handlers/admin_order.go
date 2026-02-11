package handlers

import (
	"backend/admin/backend/internal/database"
	order_entity "backend/internal/orders/entity"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)



// All Orders

func GetAllOrders(c *gin.Context) {

	var orders []order_entity.Order

	if err := database.DB.Preload("User").Preload("OrderItems").Find(&orders).Error; err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(200, orders)

}

//Update Order Status

func UpdateOrderStatus(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		response.BadRequest(c, "Invalid request")
		return
	}

	var order order_entity.Order

	if err := database.DB.First(&order, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order Not Found"})
		return
	}

	order.Status = body.Status
	database.DB.Save(&order)

	c.JSON(200, gin.H{
		"id":     order.ID,
		"status": body.Status,
	})

}

func GetOrderDetails(c *gin.Context){

 id := c.Param("id")
 
 var order order_entity.Order


 err:=database.DB.Preload("OrderItems").Preload("OrderItems.Product").First(&order,id).Error
  
 if err!=nil{
 response.BadRequest(c,"Order Not Found")
 return

}

 c.JSON(200,order)

}
