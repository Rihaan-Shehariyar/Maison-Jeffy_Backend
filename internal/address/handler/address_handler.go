package address_handler

import (
	address_entity "backend/internal/address/entity"
	address_usecase "backend/internal/address/usecase"
	"backend/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)


// Control Layer
type AddressHandler struct {
	usecase *address_usecase.AddressUsecase
}

// Dependency Injection
func NewAddressHandler(usecase *address_usecase.AddressUsecase) *AddressHandler {
	return &AddressHandler{usecase}
}


// Create Address
func (h *AddressHandler) Create(c *gin.Context) {

	userID := c.GetUint("user_id")

	var address address_entity.Address

	if err := c.ShouldBindJSON(&address); err != nil {
		response.BadRequest(c, "Invalid Json")
		return
	}

	if err := h.usecase.Create(userID, &address); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, address)

}

// Get Address

func (h *AddressHandler) GetMyAddress(c *gin.Context) {

	userID := c.GetUint("user_id")

	address, err := h.usecase.GetByUser(userID)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	c.JSON(200, address)

}


// Update Address


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


// Delete Address

func (h *AddressHandler) Delete(c *gin.Context) {

	userID := c.GetUint("user_id")
	addressID, _ := strconv.Atoi(c.Param("user_id"))

	if err := h.usecase.Delete(userID, uint(addressID)); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	c.JSON(200, gin.H{"message": "Address Deleted"})

}
