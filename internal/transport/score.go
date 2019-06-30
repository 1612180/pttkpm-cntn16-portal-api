package transport

import (
	"awesome-portal-api/internal/pkg/message"
	"awesome-portal-api/internal/service"
	"awesome-portal-api/internal/storage"
	"log"

	"github.com/gin-gonic/gin"
)

type ScoreTransport struct {
	*service.ScoreService
}

func (s *ScoreTransport) Save(c *gin.Context) {
	var score storage.Score
	if err := c.ShouldBindJSON(&score); err != nil {
		log.Println(err)
		c.JSON(200, message.Create(false))
		return
	}

	if ok := s.ScoreService.Save(&score); !ok {
		c.JSON(200, message.Create(false))
		return
	}
	c.JSON(200, message.Create(true))
}
