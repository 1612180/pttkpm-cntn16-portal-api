package models

import (
	"awesome-portal-api/internal/dtos"
	"time"
)

type Program struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	ShortName string
	LongName  string
}

func (p *Program) FromRequest(request *dtos.ProgramRequest) *Program {
	p.ShortName = request.ShortName
	p.LongName = request.LongName
	return p
}

func (p *Program) ToResponse() *dtos.ProgramResponse {
	return &dtos.ProgramResponse{
		ID:        p.ID,
		ShortName: p.ShortName,
		LongName:  p.LongName,
	}
}
