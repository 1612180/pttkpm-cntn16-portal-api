package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type FacultyTransport struct {
	*service.FacultyService
}

func (f *FacultyTransport) Save(c *gin.Context) {
	var faculty storage.Faculty
	if err := c.ShouldBindJSON(&faculty); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := f.FacultyService.Save(&faculty); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
