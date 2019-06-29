package handlers

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

type FacultyHandler struct {
	*services.FacultyService
}

func (h *FacultyHandler) FetchAll(c *gin.Context) {
	responses, ok := h.FacultyService.FetchAll()
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = responses
	c.JSON(200, msg)
}

func (h *FacultyHandler) Create(c *gin.Context) {
	// get json from client
	var request dtos.FacultyRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	// create by service
	if ok := h.FacultyService.Create(&request); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
