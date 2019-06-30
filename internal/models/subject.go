package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

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
	Semester      int  // hoc ky 1 - 2
	Status        bool // tinh trang lock hay unlock
	Weekday       int  // 2 - thu 2, 3 - thu 3
	FromPeriod    int  // tiet bat dau, 1 - 12
	ToPeriod      int  // tiet ket thuc
	ProgramID     int
	FacultyID     int
	SubjectTypeID int
}

func (s *Subject) FromRequest(request *dtos.SubjectRequest) *Subject {
	s.MaHocPhan = request.MaHocPhan
	s.Name = request.Name
	s.Class = request.Class
	s.Value = request.Value
	s.MaxStudent = request.MaxStudent
	s.Year = request.Year
	s.Semester = request.Semester
	s.Status = request.Status
	s.Weekday = request.Weekday
	s.FromPeriod = request.FromPeriod
	s.ToPeriod = request.FromPeriod
	return s
}

func (s *Subject) ToResponse() *dtos.SubjectResponse {
	return &dtos.SubjectResponse{
		ID:         s.ID,
		MaHocPhan:  s.MaHocPhan,
		Name:       s.Name,
		Class:      s.Class,
		Value:      s.Value,
		MaxStudent: s.MaxStudent,
		Year:       s.Year,
		Semester:   s.Semester,
		Status:     s.Status,
		Weekday:    s.Weekday,
		FromPeriod: s.FromPeriod,
		ToPeriod:   s.ToPeriod,
	}
}
