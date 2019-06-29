package models

import "time"

type Subject struct {
	ID            int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	MaHocPhan     string
	Name          string
	Class         string
	Value         int // so tin chi
	MaxStudent    int // sinh vien toi da
	Year          int
	Semester      int
	Status        bool // tinh trang lock hay unlock
	ProgramID     int
	FacultyID     int
	SubjectTypeID int
}
