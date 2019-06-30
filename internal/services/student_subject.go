package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type StudentSubjectService struct {
	repositories.StudentSubjectRepo
}

func (s *StudentSubjectService) FetchAll() ([]*dtos.StudentSubjectResponse, bool) {
	studentSubjects, ok := s.StudentSubjectRepo.FetchAll()
	if !ok {
		return nil, false
	}

	var responses []*dtos.StudentSubjectResponse
	for _, studentSubject := range studentSubjects {
		responses = append(responses, studentSubject.ToResponse())
	}
	return responses, true
}

func (s *StudentSubjectService) FindByID(studentID, subjectID int) (*dtos.StudentSubjectResponse, bool) {
	studentSubject, ok := s.StudentSubjectRepo.FindByID(studentID, subjectID)
	if !ok {
		return nil, false
	}
	return studentSubject.ToResponse(), true
}

func (s *StudentSubjectService) Create(request *dtos.StudentSubjectRequest) bool {
	studentSubject := (&models.StudentSubject{}).FromRequest(request)
	return s.StudentSubjectRepo.Create(studentSubject)
}
