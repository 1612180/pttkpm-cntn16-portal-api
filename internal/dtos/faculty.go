package dtos

type FacultyRequest struct {
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}

type FacultyResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}
