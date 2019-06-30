package handlers

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectHandler struct {
	*services.SubjectService
}

func (h *SubjectHandler) FetchAll(c *gin.Context) {
	responses, ok := h.SubjectService.FetchAll()
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = responses
	c.JSON(200, msg)
}

func (h *SubjectHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	response, ok := h.SubjectService.FindByID(id)
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = response
	c.JSON(200, msg)
}

func (h *SubjectHandler) Create(c *gin.Context) {
	// get json from client
	var request dtos.SubjectRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	// create by service
	if ok := h.SubjectService.Create(&request); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}

func (h *SubjectHandler) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := h.SubjectService.DeleteByID(id); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
