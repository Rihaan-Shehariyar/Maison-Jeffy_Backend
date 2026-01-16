package entity

import "time"

type OTP struct {

	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"index"`
	CodeHash  string
	ExpiresAt time.Time    
    Used bool  `gorm:"default:false"`
    CreatedAt time.Time

}