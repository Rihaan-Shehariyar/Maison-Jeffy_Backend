package address_repository

import address_entity "backend/internal/address/entity"

type AddressRepository interface {
	Create(address *address_entity.Address) error
	GetByUser(userID uint) ([]address_entity.Address, error)
	Update(address *address_entity.Address) error
	Delete(userID, addressID uint) error
	FindByID(orderID uint) (*address_entity.Address, error)
}
