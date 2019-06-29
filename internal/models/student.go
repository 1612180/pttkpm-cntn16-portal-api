package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

type Student struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string
	MSSV      string
	Year      int
	AccountID int
	ProgramID int
	FacultyID int
}

func (s *Student) FromRequest(request *dtos.StudentRequest) *Student {
	s.Name = request.Name
	s.MSSV = request.MSSV
	s.Year = request.Year
	return s
}

func (s *Student) ToResponse() (response *dtos.StudentResponse) {
	return &dtos.StudentResponse{
		ID:   s.ID,
		Name: s.Name,
		MSSV: s.MSSV,
		Year: s.Year,
	}
}
