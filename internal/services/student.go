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
	repositories.FacultyRepo
}

func (s *StudentService) FetchAll() ([]*dtos.StudentResponse, bool) {
	students, ok := s.StudentRepo.FetchAll()
	if !ok {
		return nil, false
	}

	var responses []*dtos.StudentResponse
	for _, student := range students {
		response := student.ToResponse()

		// get program
		program, ok := s.ProgramRepo.FindByID(student.ProgramID)
		if !ok {
			continue
		}
		response.ProgramShort = program.ShortName
		response.ProgramLong = program.LongName

		// get faculty
		faculty, ok := s.FacultyRepo.FindByID(student.FacultyID)
		if !ok {
			continue
		}
		response.FacultyShort = faculty.ShortName
		response.FacultyLong = faculty.LongName

		responses = append(responses, response)
	}
	return responses, true
}

func (s *StudentService) FindByID(id int) (*dtos.StudentResponse, bool) {
	student, ok := s.StudentRepo.FindByID(id)
	if !ok {
		return nil, false
	}

	response := student.ToResponse()

	// get program
	program, ok := s.ProgramRepo.FindByID(student.ProgramID)
	if !ok {
		return nil, false
	}
	response.ProgramShort = program.ShortName
	response.ProgramLong = program.LongName

	// get faculty
	faculty, ok := s.FacultyRepo.FindByID(student.FacultyID)
	if !ok {
		return nil, false
	}
	response.FacultyShort = faculty.ShortName
	response.FacultyLong = faculty.LongName

	return response, true
}

func (s *StudentService) Create(request *dtos.StudentRequest) bool {
	// create account with hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return false
	}
	account := &models.Account{
		HashedPassword: string(hashedPassword),
	}

	// create student
	student := (&models.Student{}).FromRequest(request)

	// get program
	program, ok := s.ProgramRepo.FindByShort(request.ProgramShort)
	if !ok {
		return false
	}
	student.ProgramID = program.ID

	// get faculty
	faculty, ok := s.FacultyRepo.FindByShort(request.FacultyShort)
	if !ok {
		return false
	}
	student.FacultyID = faculty.ID

	if ok := s.StudentRepo.Create(student, account); !ok {
		log.Println(err)
		return false
	}
	return true
}

func (s *StudentService) DeleteByMSSV(mssv string) bool {
	return s.StudentRepo.DeleteByMSSV(mssv)
}

func (s *StudentService) Validate(request *dtos.StudentRequest) bool {
	// find student of mssv
	student, ok := s.StudentRepo.FindByMSSV(request.MSSV)
	if !ok {
		return false
	}

	// find account of student
	account, ok := s.AccountRepo.FindByID(student.AccountID)
	if !ok {
		return false
	}

	// check password
	if err := bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(request.Password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}
