package dtos

type StudentRequest struct {
	Name         string `json:"name"`
	MSSV         string `json:"mssv"`
	Year         int    `json:"year"`
	ProgramShort string `json:"program_short"`
	FacultyShort string `json:"faculty_short"`
	Password     string `json:"password"`
}

type StudentResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	MSSV         string `json:"mssv"`
	Year         int    `json:"year"`
	ProgramShort string `json:"program_short"`
	ProgramLong  string `json:"program_long"`
	FacultyShort string `json:"faculty_short"`
	FacultyLong  string `json:"faculty_long"`
}
