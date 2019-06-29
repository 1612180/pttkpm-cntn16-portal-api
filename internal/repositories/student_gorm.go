package repositories

import (
	"awesome-portal-api/internal/models"
	"fmt"

	"github.com/jinzhu/gorm"
)

type StudentGorm struct {
	*gorm.DB
}

func (g *StudentGorm) FetchAll() ([]*models.Student, error) {
	var students []*models.Student
	if err := g.DB.Find(&students).Error; err != nil {
		return nil, err
	}
	return students, nil
}

func (g *StudentGorm) FindByID(id int) (*models.Student, error) {
	var student models.Student
	if err := g.DB.Where("id = ?", id).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (g *StudentGorm) FindByMSSV(mssv string) (*models.Student, error) {
	var student models.Student
	if err := g.DB.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		return nil, err
	}
	return &student, nil
}

func (g *StudentGorm) Create(student *models.Student, account *models.Account) error {
	tx := g.DB.Begin()

	// check if db already has mssv
	if !tx.Where("mssv = ?", student.MSSV).First(&models.Student{}).RecordNotFound() {
		tx.Rollback()
		return fmt.Errorf("mssv already exist")
	}

	// create account
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		return err
	}

	// create student
	student.AccountID = account.ID
	if err := tx.Create(student).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (g *StudentGorm) UpdateInfo(student *models.Student) error {
	return nil
}

func (g *StudentGorm) UpdatePassword(account *models.Account) error {
	return nil
}

func (g *StudentGorm) Delete(mssv string) error {
	return nil
}
