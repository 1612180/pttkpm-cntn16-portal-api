package services

import "awesome-portal-api/internal/repositories"

type Services interface {
	CreateAll() (*StudentService, *ProgramService)
}

type MyServices struct {
	repositories.StudentRepo
	repositories.AccountRepo
	repositories.ProgramRepo
}

func NewMyServices(studentRepo repositories.StudentRepo,
	accountRepo repositories.AccountRepo,
	programRepo repositories.ProgramRepo) Services {
	return &MyServices{
		StudentRepo: studentRepo,
		AccountRepo: accountRepo,
		ProgramRepo: programRepo,
	}
}

func (s *MyServices) CreateAll() (*StudentService, *ProgramService) {
	return &StudentService{StudentRepo: s.StudentRepo, AccountRepo: s.AccountRepo, ProgramRepo: s.ProgramRepo},
		&ProgramService{ProgramRepo: s.ProgramRepo}
}
