package models

import "time"

type Program struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string
	LongName  string
}
