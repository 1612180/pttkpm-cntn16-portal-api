package service

import (
	"awesome-portal-api/internal/storage"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type StudentService struct {
	storage.StudentStorage
	storage.ProgramStorage
	storage.FacultyStorage
	storage.SubjectStorage
	storage.TypeSubStorage
	storage.EnrollStorage
	storage.ScoreStorage
}

type Result struct {
	*storage.Subject `json:"subject"`
	*storage.Score   `json:"score"`
}

type StudentMore struct {
	*storage.Student `json:"student"`
	Results          []*Result `json:"results"`
}

func (s *StudentService) StudentByMSSV(mssv string) (*StudentMore, bool) {
	student, ok := s.StudentStorage.StudentByMSSV(mssv)
	if !ok {
		return nil, false
	}
	storage.FillStudent(student, s.ProgramStorage, s.FacultyStorage)

	var studentMore StudentMore
	studentMore.Student = student

	// fill results
	enrolls, ok := s.EnrollStorage.Enrolls(student.ID)
	if !ok {
		return &studentMore, true
	}

	var results []*Result
	for _, enroll := range enrolls {
		subject, ok := s.SubjectStorage.Subject(enroll.SubjectID)
		if !ok {
			continue
		}
		storage.FillSubject(subject, s.ProgramStorage, s.FacultyStorage, s.TypeSubStorage)

		score, ok := s.ScoreStorage.ScoreByEnrollID(enroll.ID)
		if !ok {
			continue
		}
		results = append(results, &Result{Subject: subject, Score: score})
	}
	studentMore.Results = results
	return &studentMore, true
}

func (s *StudentService) Save(student *storage.Student) bool {
	// create account with hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(student.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	account := &storage.Account{HashedPassword: string(hashedPassword)}

	storage.FillStudent(student, s.ProgramStorage, s.FacultyStorage)
	return s.StudentStorage.Save(student, account)
}

func (s *StudentService) DeleteByMSSV(mssv string) bool {
	return s.StudentStorage.DeleteByMSSV(mssv)
}

func (s *StudentService) Validate(student *storage.Student) bool {
	account, ok := s.StudentStorage.AccountByMSSV(student.MSSV)
	if !ok {
		return false
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(student.Password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}
