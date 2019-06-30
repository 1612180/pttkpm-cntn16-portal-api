package services

import (
	"awesome-portal-api/internal/dtos"
	"awesome-portal-api/internal/models"
	"awesome-portal-api/internal/repositories"
)

type SubjectService struct {
	repositories.SubjectRepo
	repositories.SubjectTypeRepo
	repositories.ProgramRepo
	repositories.FacultyRepo
}

func (s *SubjectService) FetchAll() ([]*dtos.SubjectResponse, bool) {
	subjects, ok := s.SubjectRepo.FetchAll()
	if !ok {
		return nil, false
	}

	var responses []*dtos.SubjectResponse
	for _, subject := range subjects {
		response := subject.ToResponse()

		// get program
		program, ok := s.ProgramRepo.FindByID(subject.ProgramID)
		if !ok {
			continue
		}
		response.ProgramShort = program.ShortName
		response.ProgramLong = program.LongName

		// get faculty
		faculty, ok := s.FacultyRepo.FindByID(subject.FacultyID)
		if !ok {
			continue
		}
		response.FacultyShort = faculty.ShortName
		response.FacultyLong = faculty.LongName

		// get type
		subjectType, ok := s.SubjectTypeRepo.FindByID(subject.SubjectTypeID)
		if !ok {
			continue
		}
		response.SubjectTypeShort = subjectType.ShortName
		response.SubjectTypeLong = subjectType.LongName

		responses = append(responses, response)
	}
	return responses, true
}

func (s *SubjectService) FindByID(id int) (*dtos.SubjectResponse, bool) {
	subject, ok := s.SubjectRepo.FindByID(id)
	if !ok {
		return nil, false
	}

	response := subject.ToResponse()

	// get program
	program, ok := s.ProgramRepo.FindByID(subject.ProgramID)
	if !ok {
		return nil, false
	}
	response.ProgramShort = program.ShortName
	response.ProgramLong = program.LongName

	// get faculty
	faculty, ok := s.FacultyRepo.FindByID(subject.FacultyID)
	if !ok {
		return nil, false
	}
	response.FacultyShort = faculty.ShortName
	response.FacultyLong = faculty.LongName

	// get type
	subjectType, ok := s.SubjectTypeRepo.FindByID(subject.SubjectTypeID)
	if !ok {
		return nil, false
	}
	response.SubjectTypeShort = subjectType.ShortName
	response.SubjectTypeLong = subjectType.LongName

	return response, true
}

func (s *SubjectService) Create(request *dtos.SubjectRequest) bool {
	subject := (&models.Subject{}).FromRequest(request)

	// get program
	program, ok := s.ProgramRepo.FindByShort(request.ProgramShort)
	if !ok {
		return false
	}
	subject.ProgramID = program.ID

	// get faculty
	faculty, ok := s.FacultyRepo.FindByShort(request.FacultyShort)
	if !ok {
		return false
	}
	subject.FacultyID = faculty.ID

	// get type
	subjectType, ok := s.SubjectTypeRepo.FindByShort(request.SubjectTypeShort)
	if !ok {
		return false
	}
	subject.SubjectTypeID = subjectType.ID

	if ok := s.SubjectRepo.Create(subject); !ok {
		return false
	}
	return true
}

func (s *SubjectService) DeleteByID(id int) bool {
	return s.SubjectRepo.DeleteByID(id)
}
