package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type FacultyGorm struct {
	*gorm.DB
}

func (g *FacultyGorm) FetchAll() ([]*models.Faculty, bool) {
	var faculties []*models.Faculty
	if err := g.DB.Find(&faculties).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return faculties, true
}

func (g *FacultyGorm) FindByID(id int) (*models.Faculty, bool) {
	var faculty models.Faculty
	if err := g.DB.Where("id = ?", id).First(&faculty).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &faculty, true
}

func (g *FacultyGorm) FindByShort(short string) (*models.Faculty, bool) {
	var faculty models.Faculty
	if err := g.DB.Where("short_name = ?", short).First(&faculty).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &faculty, true
}

func (g *FacultyGorm) Create(program *models.Faculty) bool {
	if err := g.DB.Create(program).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *FacultyGorm) DeleteByID(id int) bool {
	if err := g.DB.Where("id = ?", id).Delete(&models.Faculty{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *FacultyGorm) DeleteByShort(short string) bool {
	if err := g.DB.Where("short_name = ?", short).Delete(&models.Faculty{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
