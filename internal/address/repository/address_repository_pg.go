package address_repository

import (
	address_entity "backend/internal/address/entity"

	"gorm.io/gorm"
)

type AddressRepositoryPg struct {
	db *gorm.DB
}


func NewAddressRepositoryPg(db *gorm.DB) AddressRepository {
	return &AddressRepositoryPg{db}
}



func (r *AddressRepositoryPg) Create(address *address_entity.Address) error {
	return r.db.Create(address).Error
}

func (r *AddressRepositoryPg) GetByUser(userID uint) ([]address_entity.Address, error) {

	var address []address_entity.Address

	err := r.db.Where("user_id = ?", userID).Find(&address).Error
	return address, err

}

func (r *AddressRepositoryPg) Update(address *address_entity.Address) error {

	return r.db.Save(address).Error

}

func (r *AddressRepositoryPg) Delete(userID, addressId uint) error {

	return r.db.Where("user_id = ? AND id = ? ", userID, addressId).
		Delete(&address_entity.Address{}).Save(&address_entity.Address{}).Error
}

func (r *AddressRepositoryPg) FindByID(addressID uint) (*address_entity.Address, error) {

	var address address_entity.Address

	if err := r.db.First(&address, addressID).Error; err != nil {
		return nil, err
	}

	return &address, nil

}
