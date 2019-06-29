package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

type Faculty struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string
	LongName  string
}

func (f *Faculty) FromRequest(request *dtos.FacultyRequest) *Faculty {
	f.ShortName = request.ShortName
	f.LongName = request.LongName
	return f
}

func (f *Faculty) ToResponse() *dtos.FacultyResponse {
	return &dtos.FacultyResponse{
		ID:        f.ID,
		ShortName: f.ShortName,
		LongName:  f.LongName,
	}
}
