package models

import "time"

type StudentSubject struct {
	ID           int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ScoreMidterm float64
	ScoreFinal   float64
	ScoreRatio   int
	StudentID    int
	SubjectID    int
}
