package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type SubjectTransport struct {
	*service.SubjectService
}

func (s *SubjectTransport) Subject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	subject, ok := s.SubjectService.Subject(id)
	if !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.CreateWithData(true, subject))
}

func (s *SubjectTransport) Save(c *gin.Context) {
	var subject storage.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := s.SubjectService.Save(&subject); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
