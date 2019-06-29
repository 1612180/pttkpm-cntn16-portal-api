package repositories

import (
	"awesome-portal-api/internal/models"

	"github.com/jinzhu/gorm"
)

type FacultyGorm struct {
	*gorm.DB
}

func (g FacultyGorm) FetchAll() ([]*models.Faculty, error) {
	var faculties []*models.Faculty
	if err := g.DB.Find(&faculties).Error; err != nil {
		return nil, err
	}
	return faculties, nil
}

func (g FacultyGorm) FindByID(id int) (*models.Faculty, error) {
	var faculty models.Faculty
	if err := g.DB.Where("id = ?", id).First(&faculty).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (g FacultyGorm) FindByShort(short string) (*models.Faculty, error) {
	var faculty models.Faculty
	if err := g.DB.Where("short_name = ?", short).First(&faculty).Error; err != nil {
		return nil, err
	}
	return &faculty, nil
}

func (g FacultyGorm) Create(program *models.Faculty) error {
	return g.DB.Create(program).Error
}

func (g FacultyGorm) DeleteByID(id int) error {
	// find faculty
	var faculty models.Faculty
	if err := g.DB.Where("id = ?", id).First(&faculty).Error; err != nil {
		return err
	}

	// delete faculty
	if err := g.DB.Delete(&faculty).Error; err != nil {
		return err
	}

	return nil
}

func (g FacultyGorm) DeleteByShort(short string) error {
	// find faculty
	var faculty models.Faculty
	if err := g.DB.Where("short_name = ?", short).First(&faculty).Error; err != nil {
		return err
	}

	// delete faculty
	if err := g.DB.Delete(&faculty).Error; err != nil {
		return err
	}

	return nil
}
