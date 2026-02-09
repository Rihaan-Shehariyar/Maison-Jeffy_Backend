package entity

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	Name       string    `json:"name"`
	Email      string    `gorm:"uniqueIndex;not null" json:"email"`
	Password   string    `gorm:"not null" json:"-"` 
	Role       string    `gorm:"default:user" json:"role"`
	IsVerified bool      `gorm:"default:false" json:"is_verified"`
	IsBlocked  bool      `gorm:"default:false" json:"is_blocked"`
	CreatedAt  time.Time `json:"created_at"`
}
