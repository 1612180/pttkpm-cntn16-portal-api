package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type AccountGorm struct {
	*gorm.DB
}

func (g *AccountGorm) FindByID(id int) (*models.Account, bool) {
	var account models.Account
	if err := g.DB.Where("id = ?", id).First(&account).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &account, true
}
