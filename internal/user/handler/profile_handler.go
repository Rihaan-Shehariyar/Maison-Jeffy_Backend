package profile_handler

import (
	profile_usecase "backend/internal/user/usecase"
	"backend/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	usecase *profile_usecase.ProfileUseCase
}

 func NewProfileGHandler(u *profile_usecase.ProfileUseCase) *ProfileHandler{
    return &ProfileHandler{u}
}


func(h *ProfileHandler)GetProfile(c *gin.Context){

  userID := c.GetUint("user_id")

 user,err:= h.usecase.GetProfile(userID)
  if err!=nil{
    c.JSON(http.StatusNotFound,gin.H{"error":err.Error()})
    return
}

 c.JSON(200,gin.H{
   "id" : user.ID,
   "name" : user.Name,
   "email" : user.Email,
   
})

}

 func(h *ProfileHandler)UpdateName(c *gin.Context){
   userId:=c.GetUint("user_id")

 var req struct{
   Name string `json:"name"`
}

 if err:=c.ShouldBindJSON(&req);err!=nil{
   response.BadRequest(c,"Invalid JSON")
   return
}

 if err:= h.usecase.UpdateName(userId,req.Name);err!=nil{
    response.BadRequest(c,err.Error())
     return
}

 c.JSON(200,gin.H{"message":"Profile Name Updated"})

}

