package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
	"log"
)

type ProgramService struct {
	repositories.ProgramRepo
}

func (s *ProgramService) FetchAll() ([]*dtos.ProgramResponse, bool) {
	programs, err := s.ProgramRepo.FetchAll()
	if err != nil {
		log.Println(err)
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
	if err := s.ProgramRepo.Create(program); err != nil {
		log.Println(err)
		return false
	}
	return true
}
