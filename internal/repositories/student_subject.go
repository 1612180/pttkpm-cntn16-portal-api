package repositories

import (
	"awesome-portal-api/internal/models"
	"log"

	"github.com/jinzhu/gorm"
)

type StudentSubjectGorm struct {
	*gorm.DB
}

func (g *StudentSubjectGorm) FetchAll() ([]*models.StudentSubject, bool) {
	var studentSubjects []*models.StudentSubject
	if err := g.DB.Find(&studentSubjects).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return studentSubjects, true
}

func (g *StudentSubjectGorm) FindByID(studentID, subjectID int) (*models.StudentSubject, bool) {
	var studentSubject models.StudentSubject
	if err := g.DB.Where("student_id = ? AND subject_id = ?", studentID, subjectID).
		First(&studentSubject).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &studentSubject, true
}

func (g *StudentSubjectGorm) Create(studentSubject *models.StudentSubject) bool {
	if err := g.DB.Create(studentSubject).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
