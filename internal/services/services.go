package services

import "awesome-portal-api/internal/repositories"

type Services interface {
	CreateAll() (*StudentService, *ProgramService, *FacultyService, *SubjectService)
}

type MyServices struct {
	repositories.StudentRepo
	repositories.AccountRepo
	repositories.ProgramRepo
	repositories.FacultyRepo
	repositories.SubjectRepo
}

func NewMyServices(
	studentRepo repositories.StudentRepo,
	accountRepo repositories.AccountRepo,
	programRepo repositories.ProgramRepo,
	facultyRepo repositories.FacultyRepo,
	subjectRepo repositories.SubjectRepo,
) Services {
	return &MyServices{
		StudentRepo: studentRepo,
		AccountRepo: accountRepo,
		ProgramRepo: programRepo,
		FacultyRepo: facultyRepo,
		SubjectRepo: subjectRepo,
	}
}

func (s *MyServices) CreateAll() (*StudentService, *ProgramService, *FacultyService, *SubjectService) {
	return &StudentService{
			StudentRepo: s.StudentRepo,
			AccountRepo: s.AccountRepo,
			ProgramRepo: s.ProgramRepo,
			FacultyRepo: s.FacultyRepo,
		},
		&ProgramService{
			ProgramRepo: s.ProgramRepo,
		},
		&FacultyService{
			FacultyRepo: s.FacultyRepo,
		},
		&SubjectService{
			SubjectRepo:     s.SubjectRepo,
			SubjectTypeRepo: nil,
			ProgramRepo:     s.ProgramRepo,
			FacultyRepo:     s.FacultyRepo,
		}
}
