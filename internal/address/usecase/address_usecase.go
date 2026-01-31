package address_usecase

import (
	address_entity "backend/internal/address/entity"
	address_repository "backend/internal/address/repository"
	"errors"
)

type AddressUsecase struct {
	repo address_repository.AddressRepository
}

func NewAddressUsecase(repo address_repository.AddressRepository) *AddressUsecase {
	return &AddressUsecase{repo}
}

func (u *AddressUsecase) Create(userID uint, address *address_entity.Address) error {

	address.UserID = userID
	return u.repo.Create(address)

}

func (u *AddressUsecase) GetByUser(userID uint) ([]address_entity.Address, error) {
	return u.repo.GetByUser(userID)
}

func (u *AddressUsecase) Update(userID, addressId uint, updated *address_entity.Address) error {

	address, err := u.repo.FindByID(addressId)
	if err != nil || userID != address.UserID {
		return errors.New("Adress Not found")
	}

	address.Name = updated.Name
	address.Area = updated.Area
	address.City = updated.City
	address.State = updated.State
	address.Pincode = updated.Pincode

	return u.repo.Update(address)

}

func (u *AddressUsecase) Delete(userID, addressID uint) error {
	return u.repo.Delete(userID, addressID)
}

// func(u *AddressUsecase)FindByID(addressID uint)(*address_entity.Address,error){
//   return u.repo.FindByID(addressID)
// }
