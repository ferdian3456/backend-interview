package model

import "time"

type Transaction struct {
	ID                string
	OrderID           string
	TransactionMethod string
	Amount            float64
	Status            int
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
