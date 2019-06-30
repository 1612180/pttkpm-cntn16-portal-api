package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Subject struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	MHP        string `json:"mhp"`
	Class      string `json:"class"`
	Value      int    `json:"value"`
	MaxStudent int    `json:"max_student"`
	Status     bool   `json:"status"`

	Year       int `json:"year"`
	Semester   int `json:"semester"`
	Weekday    int `json:"weekday"`
	FromPeriod int `json:"from_period"`
	ToPeriod   int `json:"to_period"`

	ProgramID    int    `json:"-"`
	ProgramShort string `gorm:"-" json:"program_short"`
	ProgramLong  string `gorm:"-" json:"program_long"`

	FacultyID    int    `json:"-"`
	FacultyShort string `gorm:"-" json:"faculty_short"`
	FacultyLong  string `gorm:"-" json:"faculty_long"`

	TypeSubID    int    `json:"-"`
	TypeSubShort string `gorm:"-" json:"type_sub_short"`
	TypeSubLong  string `gorm:"-" json:"type_sub_long"`
}

type SubjectStorage interface {
	Subject(id int) (*Subject, bool)
	Save(subject *Subject) bool
}

type SubjectGorm struct {
	*gorm.DB
}

func NewSubjectStorage(db *gorm.DB) SubjectStorage {
	return &SubjectGorm{DB: db}
}

func (s *SubjectGorm) Subject(id int) (*Subject, bool) {
	var subject Subject
	if err := s.DB.Where("id = ?", id).First(&subject).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &subject, true
}

func (s *SubjectGorm) Save(subject *Subject) bool {
	if err := s.DB.Create(subject).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
