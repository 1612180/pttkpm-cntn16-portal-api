package repositories

import (
	"awesome-portal-api/internal/models"

	"github.com/jinzhu/gorm"
)

type AccountGorm struct {
	*gorm.DB
}

func (g *AccountGorm) FindByID(id int) (*models.Account, error) {
	var account models.Account
	if err := g.DB.Where("id = ?", id).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}
