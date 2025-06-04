package model

import "time"

type Product struct {
	ID          string
	UserID      string
	Name        string
	Description string
	Price       float64
	Available   string
	Quantity    int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
