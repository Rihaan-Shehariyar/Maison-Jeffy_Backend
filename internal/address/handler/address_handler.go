package address_handler

import (
	address_entity "backend/internal/address/entity"
	address_usecase "backend/internal/address/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AddressHandler struct {
	usecase *address_usecase.AddressUsecase
}

func NewAddressHandler(usecase *address_usecase.AddressUsecase) *AddressHandler {
	return &AddressHandler{usecase}
}

type address struct {
	UserID  uint
	Name    string `json:"name"`
	Area    string `json:"area"`
	City    string `json:"city"`
	State   string `json:"state"`
	Pincode string `json:"pincode"`
}

func (h *AddressHandler) Create(c *gin.Context) {

	userID := c.GetUint("user_id")

	var req address

	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "Invalid Json")
		return
	}

	if err := h.usecase.Create(userID, &req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, req)

}

func (h *AddressHandler) GetMyAddress(c *gin.Context) {

	userID := c.GetUint("user_id")

	address, err := h.usecase.GetByUser(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(200, address)

}

func (h *AddressHandler) Update(c *gin.Context) {

	userId := c.GetUint("user_id")
	addressId, _ := strconv.Atoi(c.Param("id"))

	var address address_entity.Address

	if err := c.ShouldBindJSON(&address); err != nil {
		response.BadRequest(c, "Invalid JSon")
		return
	}

	if err := h.usecase.Update(userId, uint(addressId), &address); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "adress Updated Succesfully"})

}


func (h *AddressHandler)Delete(c *gin.Context){
    
    userID := c.GetUint("user_id")
    addressID,_:=strconv.Atoi(c.Param("user_id"))

 if err:=h.usecase.Delete(userID,uint(addressID));err!=nil{
 response.BadRequest(c,err.Error())
 return
}

 c.JSON(200,gin.H{"message":"Address Deleted"})
 
}

f