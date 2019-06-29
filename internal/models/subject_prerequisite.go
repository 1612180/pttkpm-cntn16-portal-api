package models

import "time"

type SubjectPrerequisite struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	PreSubjectID int
	CurSubjectID int
}
