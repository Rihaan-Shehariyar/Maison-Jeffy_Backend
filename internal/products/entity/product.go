package entity

import "time"

type Product struct {

	ID          uint
	Name        string
	Description string
	Stock       int
	Price       float64
	CreatedAt   time.Time
}
