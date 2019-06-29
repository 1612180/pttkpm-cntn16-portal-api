package dtos

type FacultyRequest struct {
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}

type FacultyResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}
