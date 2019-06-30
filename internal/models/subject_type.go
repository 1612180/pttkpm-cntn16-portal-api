package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

type SubjectType struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string // bb - tc
	LongName  string // bat buoc - tu chon
}

func (s *SubjectType) FromRequest(request *dtos.SubjectTypeRequest) *SubjectType {
	s.ShortName = request.ShortName
	s.LongName = request.LongName
	return s
}

func (s *SubjectType) ToResponse() *dtos.SubjectTypeResponse {
	return &dtos.SubjectTypeResponse{
		ID:        s.ID,
		ShortName: s.ShortName,
		LongName:  s.LongName,
	}
}
