package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type SubjectTypeService struct {
	repositories.SubjectTypeRepo
}

func (s *SubjectTypeService) FetchAll() ([]*dtos.SubjectTypeResponse, bool) {
	subjectTypes, ok := s.SubjectTypeRepo.FetchAll()
	if !ok {
		return nil, false
	}

	var responses []*dtos.SubjectTypeResponse
	for _, subjectType := range subjectTypes {
		responses = append(responses, subjectType.ToResponse())
	}
	return responses, true
}

func (s *SubjectTypeService) Create(request *dtos.SubjectTypeRequest) bool {
	subjectType := (&models.SubjectType{}).FromRequest(request)
	return s.SubjectTypeRepo.Create(subjectType)
}
