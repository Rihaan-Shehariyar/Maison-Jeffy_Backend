package admin_handler

import (
	admin_usecase "backend/internal/admin/usecase"
	entitys "backend/internal/product/entity"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductAdminHandler struct {
	usecase *admin_usecase.ProductAdminUsecase
}

func NewProductAdminHandler(u *admin_usecase.ProductAdminUsecase)*ProductAdminHandler{
    return &ProductAdminHandler{u}
}

// Create Products By admin

func(h *ProductAdminHandler)CreateProduct(c *gin.Context){

 var product entitys.Product

 if err:=c.ShouldBindJSON(&product);err!=nil{
    response.BadRequest(c,"Invalid Json")
    return
}

 if err:= h.usecase.CreateProduct(&product);err!=nil{
  response.InternalError(c,"Failed To Create Product")
  return
}

c.JSON(200,product)

}

func (h *ProductAdminHandler)UpdateProduct(c *gin.Context){


 id,_ := strconv.Atoi(c.Param("id"))

var req struct{
 
 Name string `json:"name" binding:"required"`
 Description string `json:"description" binding:"required,min=10"`
 Stock int `json:"stock" binding:"required"`
 Price float64 `json:"Price" binding:"required"`
 
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
   response.BadRequest(c,"Invalid Json")
   return
}
 
 if err:= h.usecase.UpdateProduct(uint(id),req.Name,req.Description,req.Stock,req.Price);err!=nil{
    response.BadRequest(c,err.Error())
    return
}

 c.JSON(200,gin.H{"message":"Product Updated"})

}

func(h *ProductAdminHandler)DeleteProduct(c *gin.Context){
 
  id,_ := strconv.Atoi(c.Param("id"))
 
if err:=h.usecase.DeleteProduct(uint(id));err!=nil{
   response.BadRequest(c,err.Error())
   return
}
 
 c.JSON(200,gin.H{"message":"Product Deleted"})

}