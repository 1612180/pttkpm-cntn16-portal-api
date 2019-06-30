package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type TypeSubTransport struct {
	*service.TypeSubService
}

func (t *TypeSubTransport) Save(c *gin.Context) {
	var typeSub storage.TypeSub
	if err := c.ShouldBindJSON(&typeSub); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := t.TypeSubService.Save(&typeSub); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
