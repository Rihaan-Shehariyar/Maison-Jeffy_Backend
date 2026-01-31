package address_entity

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	UserID    uint   `gorm:"not null;index" json:"-"`
	Name      string `gorm:"not null" json:"name"`
	Area      string `gorm:"not null" json:"area"`
	City      string `gorm:"not null" json:"city"`
	State     string `gorm:"not null" json:"state"`
	Pincode   string `gorm:"not null" json:"pincode"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
