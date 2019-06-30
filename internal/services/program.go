package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type ProgramService struct {
	repositories.ProgramRepo
}

func (s *ProgramService) FetchAll() ([]*dtos.ProgramResponse, bool) {
	programs, ok := s.ProgramRepo.FetchAll()
	if !ok {
		return nil, false
	}

	var responses []*dtos.ProgramResponse
	for _, program := range programs {
		responses = append(responses, program.ToResponse())
	}
	return responses, true
}

func (s *ProgramService) Create(request *dtos.ProgramRequest) bool {
	program := (&models.Program{}).FromRequest(request)
	return s.ProgramRepo.Create(program)
}
