package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Subject struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MHP          string `json:"mhp"`
	Class        string `json:"class"`
	Value        int    `json:"value"`
	MaxStudent   int    `json:"max_student"`
	CountStudent int    `gorm:"-" json:"count_student"`
	Status       bool   `json:"status"`

	Year       int `json:"year"`
	Semester   int `json:"semester"`
	Weekday    int `json:"weekday"`
	FromPeriod int `json:"from_period"`
	ToPeriod   int `json:"to_period"`

	ProgramID    int    `json:"-"`
	ProgramShort string `gorm:"-" json:"program_short"`
	ProgramLong  string `gorm:"-" json:"program_long"`

	// 0 - cung chuong trinh
	// 1 - chuong trinh nao cung dang ky duoc
	CanEnroll int `json:"can_enroll"`

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
	CanTryEnroll(studentID int) ([]*Subject, bool)
	NotTryEnroll(studentID int) ([]*Subject, bool)
	CountTryEnroll(subjectID int) int
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

func (s *SubjectGorm) CanTryEnroll(studentID int) ([]*Subject, bool) {
	// find student
	var student Student
	if err := s.DB.Where("id = ?", studentID).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	// find subject in program, faculty
	var subjects []*Subject
	if err := s.DB.Where("program_id = ? AND faculty_id = ?", student.ProgramID, student.FacultyID).
		Or("can_enroll = ? AND faculty_id = ?", 1, student.FacultyID).
		Find(&subjects).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	// count value in try enroll
	var tryEnrolls []*TryEnroll
	value := 0
	if err := s.DB.Where("student_id = ?", student.ID).Find(&tryEnrolls).Error; err == nil {
		for _, tryEnroll := range tryEnrolls {
			var subject Subject
			if err := s.DB.Where("id = ?", tryEnroll.SubjectID).First(&subject).Error; err != nil {
				log.Println(err)
				continue
			}
			value += subject.Value
		}
	}

	var canSubjects []*Subject
	for _, subject := range subjects {
		// khong duoc ton tai trong try enroll
		if !s.DB.Where("student_id = ? AND subject_id = ?", student.ID, subject.ID).
			First(&TryEnroll{}).RecordNotFound() {
			continue
		}

		// khong duoc vuot qua so tin chi toi da cua sinh vien
		if value+subject.Value > student.MaxValue {
			continue
		}

		// kiem tra mon hoc da full chua
		var count int
		s.DB.Model(&TryEnroll{}).Where("subject_id = ?", subject.ID).Count(&count)
		if count+1 > subject.MaxStudent {
			continue
		}

		canSubjects = append(canSubjects, subject)
	}
	return canSubjects, true
}

func (s *SubjectGorm) NotTryEnroll(studentID int) ([]*Subject, bool) {
	// find student
	var student Student
	if err := s.DB.Where("id = ?", studentID).First(&student).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	// find subject in program, faculty
	var subjects []*Subject
	if err := s.DB.Where("program_id = ? AND faculty_id = ?", student.ProgramID, student.FacultyID).
		Or("can_enroll = ? AND faculty_id = ?", 1, student.FacultyID).
		Find(&subjects).Error; err != nil {
		log.Println(err)
		return nil, false
	}

	// count value in try enroll
	var tryEnrolls []*TryEnroll
	value := 0
	if err := s.DB.Where("student_id = ?", student.ID).Find(&tryEnrolls).Error; err == nil {
		for _, tryEnroll := range tryEnrolls {
			var subject Subject
			if err := s.DB.Where("id = ?", tryEnroll.SubjectID).First(&subject).Error; err != nil {
				log.Println(err)
				continue
			}
			value += subject.Value
		}
	}

	var notSubjects []*Subject
	for _, subject := range subjects {
		// khong duoc ton tai trong try enroll
		if !s.DB.Where("student_id = ? AND subject_id = ?", student.ID, subject.ID).
			First(&TryEnroll{}).RecordNotFound() {
			continue
		}

		// khong duoc vuot qua so tin chi toi da cua sinh vien
		if value+subject.Value > student.MaxValue {
			notSubjects = append(notSubjects, subject)
			continue
		}

		// kiem tra mon hoc da full chua
		var count int
		s.DB.Model(&TryEnroll{}).Where("subject_id = ?", subject.ID).Count(&count)
		if count+1 > subject.MaxStudent {
			notSubjects = append(notSubjects, subject)
			continue
		}
	}
	return notSubjects, true
}

func (s *SubjectGorm) CountTryEnroll(subjectID int) int {
	count := 0
	s.DB.Model(&TryEnroll{}).Where("subject_id = ?", subjectID).Count(&count)
	return count
}
