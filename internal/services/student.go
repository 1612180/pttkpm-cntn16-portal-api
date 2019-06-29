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
	students, err := s.StudentRepo.FetchAll()
	if err != nil {
		log.Println(err)
		return nil, false
	}

	var responses []*dtos.StudentResponse
	for _, student := range students {
		response := student.ToResponse()

		// get program
		program, err := s.ProgramRepo.FindByID(student.ProgramID)
		if err != nil {
			log.Println(err)
			log.Printf("student %d response failed\n", student.ID)
			continue
		}
		response.ProgramShort = program.ShortName
		response.ProgramLong = program.LongName

		// get faculty
		faculty, err := s.FacultyRepo.FindByID(student.FacultyID)
		if err != nil {
			log.Println(err)
			log.Printf("student %d response failed\n", student.ID)
			continue
		}
		response.FacultyShort = faculty.ShortName
		response.FacultyLong = faculty.LongName

		responses = append(responses, response)
	}
	return responses, true
}

func (s *StudentService) FindByID(id int) (*dtos.StudentResponse, bool) {
	student, err := s.StudentRepo.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, false
	}

	response := student.ToResponse()

	// get program
	program, err := s.ProgramRepo.FindByID(student.ProgramID)
	if err != nil {
		log.Println(err)
		log.Printf("student %d response failed\n", student.ID)
		return nil, false
	}
	response.ProgramShort = program.ShortName
	response.ProgramLong = program.LongName

	// get faculty
	faculty, err := s.FacultyRepo.FindByID(student.FacultyID)
	if err != nil {
		log.Println(err)
		log.Printf("student %d response failed\n", student.ID)
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
	program, err := s.ProgramRepo.FindByShort(request.ProgramShort)
	if err != nil {
		log.Println(err)
		log.Printf("program %s not found\n", request.ProgramShort)
		return false
	}
	student.ProgramID = program.ID

	// get faculty
	faculty, err := s.FacultyRepo.FindByShort(request.FacultyShort)
	if err != nil {
		log.Println(err)
		log.Printf("faculty %s not found\n")
		return false
	}
	student.FacultyID = faculty.ID

	if err := s.StudentRepo.Create(student, account); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *StudentService) Delete(mssv string) bool {
	if err := s.StudentRepo.DeleteByMSSV(mssv); err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *StudentService) Validate(request *dtos.StudentRequest) bool {
	// find student of mssv
	student, err := s.StudentRepo.FindByMSSV(request.MSSV)
	if err != nil {
		log.Println(err)
		return false
	}

	// find account of student
	account, err := s.AccountRepo.FindByID(student.AccountID)
	if err != nil {
		log.Println(err)
		return false
	}

	// check password
	if err = bcrypt.CompareHashAndPassword([]byte(account.HashedPassword), []byte(request.Password)); err != nil {
		log.Println(err)
		return false
	}
	return true
}
