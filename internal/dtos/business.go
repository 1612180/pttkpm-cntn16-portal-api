package dtos

type StudentAllResponse struct {
	ID           int              `json:"id"`
	Name         string           `json:"name"`
	MSSV         string           `json:"mssv"`
	Year         int              `json:"year"`
	ProgramShort string           `json:"program_short"`
	ProgramLong  string           `json:"program_long"`
	FacultyShort string           `json:"faculty_short"`
	FacultyLong  string           `json:"faculty_long"`
	Results      []ResultResponse `json:"results"`
}

type ResultResponse struct {
	Score   StudentSubjectResponse `json:"score"`
	Subject SubjectResponse        `json:"subject"`
}
