package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Account struct {
	ID             int
	HashedPassword string
}

type Student struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	MSSV string `json:"mssv"`
	Year int    `json:"year"`

	Password  string `gorm:"-" json:"password"`
	AccountID int    `json:"-"`

	ProgramID    int    `json:"-"`
	ProgramShort string `gorm:"-" json:"program_short"`
	ProgramLong  string `gorm:"-" json:"program_long"`

	FacultyID    int    `json:"-"`
	FacultyShort string `gorm:"-" json:"faculty_short"`
	FacultyLong  string `gorm:"-" json:"faculty_long"`
}

type StudentStorage interface {
	StudentByMSSV(mssv string) (*Student, bool)
	AccountByMSSV(mssv string) (*Account, bool)
	Save(student *Student, account *Account) bool
	DeleteByMSSV(mssv string) bool
}

type StudentGorm struct {
	*gorm.DB
}

func NewStudentStorage(db *gorm.DB) StudentStorage {
	return &StudentGorm{DB: db}
}

func (s *StudentGorm) StudentByMSSV(mssv string) (*Student, bool) {
	var student Student
	if err := s.DB.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	return &student, true
}

func (s *StudentGorm) AccountByMSSV(mssv string) (*Account, bool) {
	var student Student
	if err := s.DB.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	var account Account
	if err := s.DB.Where("id = ?", student.AccountID).First(&account).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &account, true
}

func (s *StudentGorm) Save(student *Student, account *Account) bool {
	tx := s.DB.Begin()

	// make sure student id not set
	student.ID = 0

	// check if db already has mssv
	if !tx.Where("mssv = ?", student.MSSV).First(&Student{}).RecordNotFound() {
		tx.Rollback()
		log.Printf("mssv %s already exist\n", student.MSSV)
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

func (s *StudentGorm) DeleteByMSSV(mssv string) bool {
	tx := s.DB.Begin()

	// find student
	var student Student
	if err := tx.Where("mssv = ?", mssv).First(&student).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return false
	}

	// delete account of student
	if err := tx.Where("id = ?", student.AccountID).Delete(&Account{}).Error; err != nil {
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
