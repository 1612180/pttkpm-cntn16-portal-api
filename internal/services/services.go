package services

import "awesome-portal-api/internal/repositories"

type Services interface {
	CreateAll() *StudentService
}

type MyServices struct {
	repositories.StudentRepo
	repositories.AccountRepo
}

func NewMyServices(studentRepo repositories.StudentRepo, accountRepo repositories.AccountRepo) Services {
	return &MyServices{
		StudentRepo: studentRepo,
		AccountRepo: accountRepo,
	}
}

func (s *MyServices) CreateAll() *StudentService {
	return &StudentService{StudentRepo: s.StudentRepo, AccountRepo: s.AccountRepo}
}
