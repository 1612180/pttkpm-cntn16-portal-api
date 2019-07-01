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

func (e *EnrollTransport) SaveMulti(c *gin.Context) {
	var multiEnroll service.MultiEnroll
	if err := c.ShouldBindJSON(&multiEnroll); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := e.EnrollService.SaveMulti(&multiEnroll); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}

func (e *EnrollTransport) SaveTryMulti(c *gin.Context) {
	var multiEnroll service.MultiEnroll
	if err := c.ShouldBindJSON(&multiEnroll); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := e.EnrollService.SaveTryMulti(&multiEnroll); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}

func (e *EnrollTransport) SaveRealAll(c *gin.Context) {
	if ok := e.EnrollService.SaveRealAll(); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}

func (e *EnrollTransport) DeleteTryMulti(c *gin.Context) {
	var multiEnroll service.MultiEnroll
	if err := c.ShouldBindJSON(&multiEnroll); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := e.EnrollService.DeleteTryMulti(&multiEnroll); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
