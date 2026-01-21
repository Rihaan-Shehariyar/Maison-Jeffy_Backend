package entitys

import "time"

type Product struct {

	ID          uint
	Name        string
	Description string
	Stock       int
	Price       float64
    ImageURL    string
	CreatedAt   time.Time
}
