package handlers

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentSubjectHandler struct {
	*services.StudentSubjectService
}

func (h *StudentSubjectHandler) FetchAll(c *gin.Context) {
	responses, ok := h.StudentSubjectService.FetchAll()
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = responses
	c.JSON(200, msg)
}

func (h *StudentSubjectHandler) FindByID(c *gin.Context) {
	studentID, err := strconv.Atoi(c.Param("student_id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	subjectID, err := strconv.Atoi(c.Param("subject_id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	response, ok := h.StudentSubjectService.FindByID(studentID, subjectID)
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = response
	c.JSON(200, msg)
}

func (h *StudentSubjectHandler) Create(c *gin.Context) {
	// get json from client
	var request dtos.StudentSubjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	// create by service
	if ok := h.StudentSubjectService.Create(&request); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
