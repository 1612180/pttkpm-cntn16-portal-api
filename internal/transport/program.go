package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type ProgramTransport struct {
	*service.ProgramService
}

func (p *ProgramTransport) Save(c *gin.Context) {
	var program storage.Program
	if err := c.ShouldBindJSON(&program); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := p.ProgramService.Save(&program); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
