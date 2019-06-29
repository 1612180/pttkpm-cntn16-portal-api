package services

import "awesome-portal-api/internal/repositories"

type Services interface {
	CreateAll() (*StudentService, *ProgramService, *FacultyService)
}

type MyServices struct {
	repositories.StudentRepo
	repositories.AccountRepo
	repositories.ProgramRepo
	repositories.FacultyRepo
}

func NewMyServices(studentRepo repositories.StudentRepo,
	accountRepo repositories.AccountRepo,
	programRepo repositories.ProgramRepo,
	facultyRepo repositories.FacultyRepo,
) Services {
	return &MyServices{
		StudentRepo: studentRepo,
		AccountRepo: accountRepo,
		ProgramRepo: programRepo,
		FacultyRepo: facultyRepo,
	}
}

func (s *MyServices) CreateAll() (*StudentService, *ProgramService, *FacultyService) {
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
		}
}
