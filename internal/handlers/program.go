package handlers

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

type ProgramHandler struct {
	*services.ProgramService
}

func (h *ProgramHandler) FetchAll(c *gin.Context) {
	responses, err := h.ProgramService.FetchAll()
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = responses
	c.JSON(200, msg)
}

func (h *ProgramHandler) Create(c *gin.Context) {
	// get json from client
	var request dtos.ProgramRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	// create by service
	if err := h.ProgramService.Create(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
