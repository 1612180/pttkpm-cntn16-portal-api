package handlers

import "awesome-portal-api/internal/services"

type Handlers interface {
	CreateAll() *StudentHandler
}

type MyHandlers struct {
	studentService *services.StudentService
}

func NewMyHandlers(studentService *services.StudentService) Handlers {
	return &MyHandlers{studentService: studentService}
}

func (h *MyHandlers) CreateAll() *StudentHandler {
	return &StudentHandler{studentService: h.studentService}
}
