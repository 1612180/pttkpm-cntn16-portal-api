package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type StudentTransport struct {
	*service.StudentService
}

func (s *StudentTransport) StudentByMSSV(c *gin.Context) {
	studentMore, ok := s.StudentService.StudentByMSSV(c.Param("mssv"))
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.CreateWithData(true, studentMore))
}

func (s *StudentTransport) TryEnroll(c *gin.Context) {
	if c.Param("status") == "already" {
		studentMore, ok := s.StudentService.AlreadyTryEnroll(c.Param("mssv"))
		if !ok {
			c.JSON(200, message.Create(false))
			return
		}
		c.JSON(200, message.CreateWithData(true, studentMore))
		return
	} else if c.Param("status") == "can" {
		studentMore, ok := s.StudentService.CanTryEnroll(c.Param("mssv"))
		if !ok {
			c.JSON(200, message.Create(false))
			return
		}
		c.JSON(200, message.CreateWithData(true, studentMore))
		return
	} else if c.Param("status") == "not" {
		studentMore, ok := s.StudentService.NotTryEnroll(c.Param("mssv"))
		if !ok {
			c.JSON(200, message.Create(false))
			return
		}
		c.JSON(200, message.CreateWithData(true, studentMore))
		return
	}
	c.JSON(200, message.Create(false))
}

func (s *StudentTransport) Save(c *gin.Context) {
	var student storage.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := s.StudentService.Save(&student); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}

func (s *StudentTransport) DeleteByMSSV(c *gin.Context) {
	if ok := s.StudentService.DeleteByMSSV(c.Param("mssv")); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}

func (s *StudentTransport) Validate(c *gin.Context) {
	var student storage.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := s.StudentService.Validate(&student); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
