package service

import "awesome-portal-api/internal/storage"

type SubjectService struct {
	storage.SubjectStorage
	storage.ProgramStorage
	storage.FacultyStorage
	storage.TypeSubStorage
}

func (s *SubjectService) Subject(id int) (*storage.Subject, bool) {
	subject, ok := s.SubjectStorage.Subject(id)
	if !ok {
		return nil, false
	}

	storage.FillSubject(subject, s.ProgramStorage, s.FacultyStorage, s.TypeSubStorage)
	return subject, true
}

func (s *SubjectService) Save(subject *storage.Subject) bool {
	storage.FillSubject(subject, s.ProgramStorage, s.FacultyStorage, s.TypeSubStorage)
	return s.SubjectStorage.Save(subject)
}
