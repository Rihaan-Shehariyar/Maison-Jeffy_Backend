package admin_handler

import (
	admin_usecase "backend/internal/admin/usecase"
	entitys "backend/internal/product/entity"
	"backend/pkg/response"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

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

 name:=c.PostForm("name")
 priceStr := c.PostForm("price")
 stockStr := c.PostForm("stock") 
 description := c.PostForm("description")



  if name == "" || priceStr == "" || stockStr == ""{
   response.BadRequest(c,"Name,Price and Stock is required ")
    return
}

 price,err:=strconv.ParseFloat(priceStr,64)
 if err!=nil{
   response.BadRequest(c,"Invalid Price")
   return
}
 
 stock,err := strconv.Atoi(stockStr) 
  if err!=nil || stock <0{
    response.BadRequest(c,"Invalid Stock")
    return
}

 file,err:= c.FormFile("image")
 if err!=nil{
  response.BadRequest(c,"Image is required")
  return
}

uploadDir := "uploads/products"
_ = os.Mkdir(uploadDir,os.ModePerm)


fileName := fmt.Sprintf("%d_%s",time.Now().Unix(),file.Filename)
imagePath := filepath.Join(uploadDir,fileName)

if err:= c.SaveUploadedFile(file,imagePath);err!=nil{
  response.InternalError(c,"Failed to Save Image")
  return
}
 
product := &entitys.Product{
  Name: name,
  Description: description,
  Price: price,
  Stock: stock,
  ImageURL: imagePath,
}

 if err:=h.usecase.CreateProduct(product);err!=nil{
   response.InternalError(c,"Failed To Create Products")
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