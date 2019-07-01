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

type TryEnroll struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
}

type EnrollStorage interface {
	EnrollsByStudentID(studentID int) ([]*Enroll, bool)
	TryEnrollsByStudentID(studentID int) ([]*TryEnroll, bool)
	TryEnrolls() ([]*TryEnroll, bool)
	Save(enroll *Enroll) bool
	SaveTry(tryEnroll *TryEnroll) bool
	SaveReal(tryEnroll *TryEnroll) bool
	DeleteTrySSID(studentID, subjectID int) bool
}

type EnrollGorm struct {
	*gorm.DB
}

func NewEnrollStorage(db *gorm.DB) EnrollStorage {
	return &EnrollGorm{DB: db}
}

func (e *EnrollGorm) EnrollsByStudentID(studentID int) ([]*Enroll, bool) {
	var enrolls []*Enroll
	if err := e.DB.Where("student_id = ?", studentID).Find(&enrolls).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return enrolls, true
}

func (e *EnrollGorm) TryEnrollsByStudentID(studentID int) ([]*TryEnroll, bool) {
	var tryEnrolls []*TryEnroll
	if err := e.DB.Where("student_id = ?", studentID).Find(&tryEnrolls).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return tryEnrolls, true
}

func (e *EnrollGorm) TryEnrolls() ([]*TryEnroll, bool) {
	var tryEnrolls []*TryEnroll
	if err := e.DB.Find(&tryEnrolls).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return tryEnrolls, true
}

// TODO dieu kien
func (e *EnrollGorm) Save(enroll *Enroll) bool {
	tx := e.DB.Begin()

	// already enroll
	if !tx.Where("student_id = ? AND subject_id = ?", enroll.StudentID, enroll.SubjectID).
		First(&Enroll{}).RecordNotFound() {
		log.Printf("student %d already enroll subject %d\n", enroll.StudentID, enroll.SubjectID)
		tx.Rollback()
		return false
	}

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

// TODO dieu kien
func (e *EnrollGorm) SaveTry(tryEnroll *TryEnroll) bool {
	tx := e.DB.Begin()

	// already try enroll
	if !tx.Where("student_id = ? AND subject_id = ?", tryEnroll.StudentID, tryEnroll.SubjectID).
		First(&TryEnroll{}).RecordNotFound() {
		log.Printf("student %d already try enroll subject %d\n", tryEnroll.StudentID, tryEnroll.SubjectID)
		tx.Rollback()
		return false
	}

	if err := tx.Create(tryEnroll).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}

func (e *EnrollGorm) SaveReal(tryEnroll *TryEnroll) bool {
	tx := e.DB.Begin()

	enroll := Enroll{
		StudentID: tryEnroll.StudentID,
		SubjectID: tryEnroll.SubjectID,
	}
	if err := tx.Create(&enroll).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	if err := tx.Delete(tryEnroll).Error; err != nil {
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

func (e *EnrollGorm) DeleteTrySSID(studentID, subjectID int) bool {
	tx := e.DB.Begin()

	var tryEnroll TryEnroll
	if err := e.DB.Where("student_id = ? AND subject_id = ?", studentID, subjectID).First(&tryEnroll).
		Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	if err := e.DB.Delete(&tryEnroll).Error; err != nil {
		log.Println(err)
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true
}
