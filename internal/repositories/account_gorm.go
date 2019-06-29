package repositories

import (
	"awesome-portal-api/internal/models"

	"github.com/jinzhu/gorm"
)

type AccountGorm struct {
	*gorm.DB
}

func (g *AccountGorm) FindByID(id int) (account *models.Account, err error) {
	err = g.DB.Where("id = ?", id).First(account).Error
	return
}
