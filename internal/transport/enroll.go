package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

type EnrollTransport struct {
	*service.EnrollService
}

func (e *EnrollTransport) Save(c *gin.Context) {
	var multiEnroll service.MultiEnroll
	if err := c.ShouldBindJSON(&multiEnroll); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := e.EnrollService.Save(&multiEnroll); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
