package model

import "time"

type Order struct {
	ID         string
	ProductID  string
	CustomerID string
	SellerID   string
	CreatedBy  string
	Amount     float64
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
