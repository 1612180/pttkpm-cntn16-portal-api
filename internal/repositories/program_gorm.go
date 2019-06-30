package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type ProgramGorm struct {
	*gorm.DB
}

func (g *ProgramGorm) FetchAll() ([]*models.Program, bool) {
	var programs []*models.Program
	if err := g.DB.Find(&programs).Error; err != nil {
		return nil, false
	}
	return programs, true
}

func (g *ProgramGorm) FindByID(id int) (*models.Program, bool) {
	var program models.Program
	if err := g.DB.Where("id = ?", id).First(&program).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &program, true
}

func (g *ProgramGorm) FindByShort(short string) (*models.Program, bool) {
	var program models.Program
	if err := g.DB.Where("short_name = ?", short).First(&program).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &program, true
}

func (g *ProgramGorm) Create(program *models.Program) bool {
	if err := g.DB.Create(program).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *ProgramGorm) DeleteByID(id int) bool {
	if err := g.DB.Where("id = ?", id).Delete(&models.Program{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (g *ProgramGorm) DeleteByShort(short string) bool {
	if err := g.DB.Where("short_name = ?", short).Delete(&models.Program{}).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
