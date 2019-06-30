package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Enroll struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
}

type EnrollStorage interface {
	Enrolls(studentID int) ([]*Enroll, bool)
	Save(enroll *Enroll) bool
}

type EnrollGorm struct {
	*gorm.DB
}

func NewEnrollStorage(db *gorm.DB) EnrollStorage {
	return &EnrollGorm{DB: db}
}

func (e *EnrollGorm) Enrolls(studentID int) ([]*Enroll, bool) {
	var enrolls []*Enroll
	if err := e.DB.Where("student_id = ?", studentID).Find(&enrolls).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return enrolls, true
}

// TODO check dieu kien
func (e *EnrollGorm) Save(enroll *Enroll) bool {
	tx := e.DB.Begin()

	if err := tx.Create(enroll).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	// auto create score with default 0
	if err := tx.Create(&Score{EnrollID: enroll.ID}).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}
