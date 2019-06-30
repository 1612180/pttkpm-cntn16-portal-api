package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Faculty struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type FacultyStorage interface {
	Faculty(id int) (*Faculty, bool)
	FacultyByShort(short string) (*Faculty, bool)
	Save(faculty *Faculty) bool
}

type FacultyGorm struct {
	*gorm.DB
}

func NewFacultyStorage(db *gorm.DB) FacultyStorage {
	return &FacultyGorm{DB: db}
}

func (f *FacultyGorm) Faculty(id int) (*Faculty, bool) {
	var faculty Faculty
	if err := f.DB.Where("id = ?", id).First(&faculty).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &faculty, true
}

func (f *FacultyGorm) FacultyByShort(short string) (*Faculty, bool) {
	var faculty Faculty
	if err := f.DB.Where("short = ?", short).First(&faculty).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &faculty, true
}

func (f *FacultyGorm) Save(faculty *Faculty) bool {
	if err := f.DB.Create(faculty).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
