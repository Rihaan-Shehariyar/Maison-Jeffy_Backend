package entity

import "time"

type User struct {

	ID         uint      `gorm:"primaryKey"`
	Name       string
	Email      string    `gorm:"uniqueIndex;not null"`
	Password   string    `gorm:"not null"`
	Role       string      `gorm:"default:user"`
	IsVerified bool      `gorm:"default:false"`
    IsBlocked bool       `gorm:"default:false"`
	CreatedAt  time.Time
}

