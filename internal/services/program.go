package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type ProgramService struct {
	repositories.ProgramRepo
}

func (s *ProgramService) FetchAll() ([]*dtos.ProgramResponse, error) {
	programs, err := s.ProgramRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	var responses []*dtos.ProgramResponse
	for _, program := range programs {
		responses = append(responses, program.ToResponse())
	}
	return responses, nil
}

func (s *ProgramService) Create(request *dtos.ProgramRequest) error {
	program := (&models.Program{}).FromRequest(request)
	return s.ProgramRepo.Create(program)
}
