package entity

import "time"

type User struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Role      string
	IsActive  bool
	CreatedAt time.Time
}

