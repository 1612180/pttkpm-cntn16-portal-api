package dtos

type SubjectRequest struct {
	MaHocPhan        string `json:"mahocphan"`
	Name             string `json:"name"`
	Class            string `json:"class"`
	Value            int    `json:"value"`
	MaxStudent       int    `json:"max_student"`
	Year             int    `json:"year"`
	Semester         int    `json:"semester"`
	Status           bool   `json:"status"`
	Weekday          int    `json:"weekday"`
	FromPeriod       int    `json:"from_period"`
	ToPeriod         int    `json:"to_period"`
	ProgramShort     string `json:"program_short"`
	FacultyShort     string `json:"faculty_short"`
	SubjectTypeShort string `json:"subject_type_short"`
}

type SubjectResponse struct {
	ID               int    `json:"id"`
	MaHocPhan        string `json:"mahocphan"`
	Name             string `json:"name"`
	Class            string `json:"class"`
	Value            int    `json:"value"`
	MaxStudent       int    `json:"max_student"`
	Year             int    `json:"year"`
	Semester         int    `json:"semester"`
	Status           bool   `json:"status"`
	Weekday          int    `json:"weekday"`
	FromPeriod       int    `json:"from_period"`
	ToPeriod         int    `json:"to_period"`
	ProgramShort     string `json:"program_short"`
	ProgramLong      string `json:"program_long"`
	FacultyShort     string `json:"faculty_short"`
	FacultyLong      string `json:"faculty_long"`
	SubjectTypeShort string `json:"subject_type"`
	SubjectTypeLong  string `json:"subject_type_long"`
}
