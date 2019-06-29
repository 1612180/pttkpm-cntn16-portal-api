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
	responses, err := h.StudentService.FetchAll()
	if err != nil {
		log.Println(err)
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

	response, err := h.StudentService.FindByID(id)
	if err != nil {
		log.Println(err)
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
	if err := h.StudentService.Create(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}

func (h *StudentHandler) Delete(c *gin.Context) {
	if err := h.StudentService.Delete(c.Param("mssv")); err != nil {
		log.Println(err)
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

	if err := h.StudentService.Validate(&request); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	c.JSON(200, message.Create(true))
}
