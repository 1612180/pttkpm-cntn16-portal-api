package repositories

import (
	"awesome-portal-api/internal/models"

	"github.com/jinzhu/gorm"
)

type ProgramGorm struct {
	*gorm.DB
}

func (g ProgramGorm) FetchAll() ([]*models.Program, error) {
	var programs []*models.Program
	if err := g.DB.Find(&programs).Error; err != nil {
		return nil, err
	}
	return programs, nil
}

func (g ProgramGorm) FindByID(id int) (*models.Program, error) {
	var program models.Program
	if err := g.DB.Where("id = ?", id).First(&program).Error; err != nil {
		return nil, err
	}
	return &program, nil
}

func (g ProgramGorm) FindByShort(short string) (*models.Program, error) {
	var program models.Program
	if err := g.DB.Where("short_name = ?", short).First(&program).Error; err != nil {
		return nil, err
	}
	return &program, nil
}

func (g ProgramGorm) Create(program *models.Program) error {
	return g.DB.Create(program).Error
}

func (g ProgramGorm) DeleteByID(id int) error {
	// find program
	var program models.Program
	if err := g.DB.Where("id = ?", id).First(&program).Error; err != nil {
		return err
	}

	// delete program
	if err := g.DB.Delete(&program).Error; err != nil {
		return err
	}

	return nil
}

func (g ProgramGorm) DeleteByShort(short string) error {
	// find program
	var program models.Program
	if err := g.DB.Where("short_name = ?", short).First(&program).Error; err != nil {
		return err
	}

	// delete program
	if err := g.DB.Delete(&program).Error; err != nil {
		return err
	}

	return nil
}
