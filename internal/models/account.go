package models

import "time"

type Account struct {
	ID             int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	HashedPassword string
}
