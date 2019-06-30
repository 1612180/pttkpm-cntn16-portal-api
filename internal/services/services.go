package services

import "awesome-portal-api/internal/repositories"

type Services interface {
	CreateAll() (
		*StudentService,
		*ProgramService,
		*FacultyService,
		*SubjectService,
		*SubjectTypeService,
	)
}

type MyServices struct {
	repositories.StudentRepo
	repositories.AccountRepo
	repositories.ProgramRepo
	repositories.FacultyRepo
	repositories.SubjectRepo
	repositories.SubjectTypeRepo
}

func NewMyServices(
	studentRepo repositories.StudentRepo,
	accountRepo repositories.AccountRepo,
	programRepo repositories.ProgramRepo,
	facultyRepo repositories.FacultyRepo,
	subjectRepo repositories.SubjectRepo,
	subjectTypeRepo repositories.SubjectTypeRepo,
) Services {
	return &MyServices{
		StudentRepo:     studentRepo,
		AccountRepo:     accountRepo,
		ProgramRepo:     programRepo,
		FacultyRepo:     facultyRepo,
		SubjectRepo:     subjectRepo,
		SubjectTypeRepo: subjectTypeRepo,
	}
}

func (s *MyServices) CreateAll() (
	*StudentService,
	*ProgramService,
	*FacultyService,
	*SubjectService,
	*SubjectTypeService,
) {
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
			SubjectTypeRepo: s.SubjectTypeRepo,
			ProgramRepo:     s.ProgramRepo,
			FacultyRepo:     s.FacultyRepo,
		},
		&SubjectTypeService{
			SubjectTypeRepo: s.SubjectTypeRepo,
		}
}
