package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
	"log"
)

type FacultyService struct {
	repositories.FacultyRepo
}

func (s *FacultyService) FetchAll() ([]*dtos.FacultyResponse, bool) {
	faculties, err := s.FacultyRepo.FetchAll()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	var responses []*dtos.FacultyResponse
	for _, faculty := range faculties {
		responses = append(responses, faculty.ToResponse())
	}
	return responses, true
}

func (s *FacultyService) Create(request *dtos.FacultyRequest) bool {
	faculty := (&models.Faculty{}).FromRequest(request)
	if err := s.FacultyRepo.Create(faculty); err != nil {
		log.Println(err)
		return false
	}
	return true
}
