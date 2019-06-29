package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type StudentService struct {
	repositories.StudentRepo
	repositories.AccountRepo
	repositories.ProgramRepo
}

func (s *StudentService) FetchAll() ([]*dtos.StudentResponse, error) {
	students, err := s.StudentRepo.FetchAll()
	if err != nil {
		return nil, err
	}

	var responses []*dtos.StudentResponse
	for _, student := range students {
		response := student.ToResponse()

		// get program
		program, err := s.ProgramRepo.FindByID(student.ProgramID)
		if err != nil {
			log.Println(err)
			log.Println("one response failed")
			continue
		}
		response.ProgramShort = program.ShortName
		response.ProgramLong = program.LongName

		responses = append(responses, response)
	}
	return responses, nil
}

func (s *StudentService) FindByID(id int) (*dtos.StudentResponse, error) {
	student, err := s.StudentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return student.ToResponse(), nil
}

func (s *StudentService) Create(request *dtos.StudentRequest) error {
	// create account with hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	account := &models.Account{
		HashedPassword: string(hashedPassword),
	}

	// create student
	student := (&models.Student{}).FromRequest(request)

	// get program
	program, err := s.ProgramRepo.FindByShort(request.ProgramShort)
	if err != nil {
		log.Println("program not found")
		return err
	}
	student.ProgramID = program.ID

	return s.StudentRepo.Create(student, account)
}

func (s *StudentService) Delete(mssv string) error {
	return s.StudentRepo.DeleteByMSSV(mssv)
}

func (s *StudentService) Validate(request *dtos.StudentRequest) error {
	// find student of mssv
	student, err := s.StudentRepo.FindByMSSV(request.MSSV)
	if err != nil {
		return err
	}

	// find account of student
	account, err := s.AccountRepo.FindByID(student.AccountID)
	if err != nil {
		return err
	}

	// check password
	if err = bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(request.Password)); err != nil {
		return err
	}
	return nil
}
