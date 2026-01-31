package address_entity

import "time"

type Address struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"not null;index"`
	Name      string
	Area      string
	City      string
	State     string
	Pincode   string
	CreatedAt time.Time
}
