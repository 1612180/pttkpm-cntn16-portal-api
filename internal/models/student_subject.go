package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

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

func (s *StudentSubject) FromRequest(request *dtos.StudentSubjectRequest) *StudentSubject {
	s.ScoreMidterm = request.ScoreMidterm
	s.ScoreFinal = request.ScoreFinal
	s.ScoreRatio = request.ScoreRatio
	s.StudentID = request.StudentID
	s.SubjectID = request.SubjectID
	return s
}

func (s *StudentSubject) ToResponse() *dtos.StudentSubjectResponse {
	return &dtos.StudentSubjectResponse{
		ID:           s.ID,
		ScoreMidterm: s.ScoreMidterm,
		ScoreFinal:   s.ScoreFinal,
		ScoreRatio:   s.ScoreRatio,
		StudentID:    s.StudentID,
		SubjectID:    s.SubjectID,
	}
}
