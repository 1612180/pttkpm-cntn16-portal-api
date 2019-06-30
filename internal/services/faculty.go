package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type FacultyService struct {
	repositories.FacultyRepo
}

func (s *FacultyService) FetchAll() ([]*dtos.FacultyResponse, bool) {
	faculties, ok := s.FacultyRepo.FetchAll()
	if !ok {
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
	return s.FacultyRepo.Create(faculty)
}
