package handlers

import "awesome-portal-api/internal/services"

type Handlers interface {
	CreateAll() (
		*StudentHandler,
		*ProgramHandler,
		*FacultyHandler,
		*SubjectHandler,
		*SubjectTypeHandler,
	)
}

type MyHandlers struct {
	*services.StudentService
	*services.ProgramService
	*services.FacultyService
	*services.SubjectService
	*services.SubjectTypeService
}

func NewMyHandlers(studentService *services.StudentService,
	programService *services.ProgramService,
	facultyService *services.FacultyService,
	subjectService *services.SubjectService,
	subjectTypeService *services.SubjectTypeService,
) Handlers {
	return &MyHandlers{
		StudentService:     studentService,
		ProgramService:     programService,
		FacultyService:     facultyService,
		SubjectService:     subjectService,
		SubjectTypeService: subjectTypeService,
	}
}

func (h *MyHandlers) CreateAll() (
	*StudentHandler,
	*ProgramHandler,
	*FacultyHandler,
	*SubjectHandler,
	*SubjectTypeHandler,
) {
	return &StudentHandler{
			StudentService: h.StudentService,
		},
		&ProgramHandler{
			ProgramService: h.ProgramService,
		},
		&FacultyHandler{
			FacultyService: h.FacultyService,
		},
		&SubjectHandler{
			SubjectService: h.SubjectService,
		},
		&SubjectTypeHandler{
			SubjectTypeService: h.SubjectTypeService,
		}
}
