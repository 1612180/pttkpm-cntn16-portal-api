package services

import "awesome-portal-api/internal/repositories"

type BusinessService struct {
	repositories.StudentRepo
	repositories.ProgramRepo
	repositories.FacultyRepo
	repositories.SubjectRepo
	repositories.SubjectTypeRepo
}
