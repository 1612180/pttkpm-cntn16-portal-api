package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type SubjectTypeGorm struct {
	*gorm.DB
}

func (g *SubjectTypeGorm) FetchAll() ([]*models.SubjectType, bool) {
	var subjectTypes []*models.SubjectType
	if err := g.DB.Find(&subjectTypes).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return subjectTypes, true
}

func (g *SubjectTypeGorm) FindByID(id int) (*models.SubjectType, bool) {
	var subjectType models.SubjectType
	if err := g.DB.Where("id = ?", id).First(&subjectType).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &subjectType, true
}

func (g *SubjectTypeGorm) FindByShort(short string) (*models.SubjectType, bool) {
	var subjectType models.SubjectType
	if err := g.DB.Where("short_name = ?", short).First(&subjectType).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &subjectType, true
}

func (g *SubjectTypeGorm) Create(subjectType *models.SubjectType) bool {
	if err := g.DB.Create(subjectType).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *SubjectTypeGorm) DeleteByID(id int) bool {
	if err := g.DB.Where("id = ?", id).Delete(&models.SubjectType{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *SubjectTypeGorm) DeleteByShort(short string) bool {
	if err := g.DB.Where("short_name = ?", short).Delete(&models.SubjectType{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
