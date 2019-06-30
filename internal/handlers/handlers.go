package handlers

import "awesome-portal-api/internal/services"

type Handlers interface {
	CreateAll() (
		*StudentHandler,
		*ProgramHandler,
		*FacultyHandler,
		*SubjectHandler,
	)
}

type MyHandlers struct {
	*services.StudentService
	*services.ProgramService
	*services.FacultyService
	*services.SubjectService
}

func NewMyHandlers(studentService *services.StudentService,
	programService *services.ProgramService,
	facultyService *services.FacultyService,
	subjectService *services.SubjectService,
) Handlers {
	return &MyHandlers{
		StudentService: studentService,
		ProgramService: programService,
		FacultyService: facultyService,
		SubjectService: subjectService,
	}
}

func (h *MyHandlers) CreateAll() (
	*StudentHandler,
	*ProgramHandler,
	*FacultyHandler,
	*SubjectHandler,
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
		}
}
