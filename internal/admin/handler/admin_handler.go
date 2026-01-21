package admin_handler

import (
	admin_usecase "backend/internal/admin/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserAdminHandler struct {
	usecase *admin_usecase.UserAdminUsecase
}

func NewUserAdminHandler(u *admin_usecase.UserAdminUsecase) *UserAdminHandler{
    return &UserAdminHandler{u}
}

func(h *UserAdminHandler)UpdateUser(c *gin.Context){
 
 id,_ := strconv.Atoi(c.Param("id"))

 var req struct{
   Name string `json:"name"`
   Email string `json:"email" binding:"required,email"`
   Role string `json:"role" binding:"required"`
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
   response.BadRequest(c,"Invalid Json")
   return
}

 if err:= h.usecase.UpdateUser(uint(id),req.Name,req.Email,req.Role);err!=nil{
   response.BadRequest(c,err.Error())
   return
}

c.JSON(200,gin.H{"message":"Updated User Successfully"})
 
}

func(h *UserAdminHandler)BlockUser(c *gin.Context){

  id,_:=strconv.Atoi(c.Param("id")) 

  if err:=h.usecase.BlockUser(uint(id));err!=nil{
    response.BadRequest(c,err.Error())
    return
}

 c.JSON(200,gin.H{"message":"User Blocked"})

}
