package handlers

import "awesome-portal-api/internal/services"

type Handlers interface {
	CreateAll() (*StudentHandler, *ProgramHandler, *FacultyHandler)
}

type MyHandlers struct {
	*services.StudentService
	*services.ProgramService
	*services.FacultyService
}

func NewMyHandlers(studentService *services.StudentService,
	programService *services.ProgramService,
	facultyService *services.FacultyService,
) Handlers {
	return &MyHandlers{
		StudentService: studentService,
		ProgramService: programService,
		FacultyService: facultyService,
	}
}

func (h *MyHandlers) CreateAll() (*StudentHandler, *ProgramHandler, *FacultyHandler) {
	return &StudentHandler{
			StudentService: h.StudentService,
		},
		&ProgramHandler{
			ProgramService: h.ProgramService,
		},
		&FacultyHandler{
			FacultyService: h.FacultyService,
		}
}
