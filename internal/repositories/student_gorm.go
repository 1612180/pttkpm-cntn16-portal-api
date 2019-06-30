package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type StudentGorm struct {
	*gorm.DB
}

func (g *StudentGorm) FetchAll() ([]*models.Student, bool) {
	var students []*models.Student
	if err := g.DB.Find(&students).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return students, true
}

func (g *StudentGorm) FindByID(id int) (*models.Student, bool) {
	var student models.Student
	if err := g.DB.Where("id = ?", id).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &student, true
}

func (g *StudentGorm) FindByMSSV(mssv string) (*models.Student, bool) {
	var student models.Student
	if err := g.DB.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &student, true
}

func (g *StudentGorm) Create(student *models.Student, account *models.Account) bool {
	tx := g.DB.Begin()

	// check if db already has mssv
	if !tx.Where("mssv = ?", student.MSSV).First(&models.Student{}).RecordNotFound() {
		tx.Rollback()
		return false
	}

	// create account
	if err := tx.Create(account).Error; err != nil {
		tx.Rollback()
		log.Print(err)
		return false
	}

	// create student
	student.AccountID = account.ID
	if err := tx.Create(student).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	tx.Commit()
	return true
}

func (g *StudentGorm) DeleteByMSSV(mssv string) bool {
	tx := g.DB.Begin()

	// find student
	var student models.Student
	if err := tx.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	// delete account of student
	if err := tx.Where("id = ?", student.AccountID).Delete(&models.Account{}).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	// delete student
	if err := tx.Delete(&student).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	tx.Commit()
	return true
}
