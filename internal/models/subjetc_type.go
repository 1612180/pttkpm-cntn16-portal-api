package models

import "time"

type SubjectType struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string // bb - tc
	LongName  string // bat buoc - tu chon
}
