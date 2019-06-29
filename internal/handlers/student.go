package handlers

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/services"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type StudentHandler struct {
	*services.StudentService
}

func (h *StudentHandler) FetchAll(c *gin.Context) {
	responses, ok := h.StudentService.FetchAll()
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = responses
	c.JSON(200, msg)
}

func (h *StudentHandler) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	response, ok := h.StudentService.FindByID(id)
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}

	msg := message.Create(true)
	msg["data"] = response
	c.JSON(200, msg)
}

func (h *StudentHandler) Create(c *gin.Context) {
	// get json from client
	var request dtos.StudentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	// create by service
	if ok := h.StudentService.Create(&request); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}

func (h *StudentHandler) Delete(c *gin.Context) {
	if ok := h.StudentService.Delete(c.Param("mssv")); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}

func (h *StudentHandler) Login(c *gin.Context) {
	// get json from client
	var request dtos.StudentRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := h.StudentService.Validate(&request); !ok {
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
