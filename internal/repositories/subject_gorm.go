package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type SubjectGorm struct {
	*gorm.DB
}

func (g *SubjectGorm) FetchAll() ([]*models.Subject, bool) {
	var subjects []*models.Subject
	if err := g.DB.Find(&subjects).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return subjects, true
}

func (g *SubjectGorm) FindByID(id int) (*models.Subject, bool) {
	var subject models.Subject
	if err := g.DB.Where("id = ?", id).First(&subject).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &subject, true
}

func (g *SubjectGorm) Create(subject *models.Subject) bool {
	if err := g.DB.Create(subject).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *SubjectGorm) DeleteByID(id int) bool {
	if err := g.DB.Where("id = ?", id).Delete(&models.Subject{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
