package model

import "time"

type User struct {
	ID           string
	Username     string
	Email        string
	ContactPhone string
	Password     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
