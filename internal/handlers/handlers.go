package handlers

import "awesome-portal-api/internal/services"

type Handlers interface {
	CreateAll() (*StudentHandler, *ProgramHandler)
}

type MyHandlers struct {
	*services.StudentService
	*services.ProgramService
}

func NewMyHandlers(studentService *services.StudentService,
	programService *services.ProgramService) Handlers {
	return &MyHandlers{StudentService: studentService, ProgramService: programService}
}

func (h *MyHandlers) CreateAll() (*StudentHandler, *ProgramHandler) {
	return &StudentHandler{StudentService: h.StudentService},
		&ProgramHandler{ProgramService: h.ProgramService}
}
