package dtos

type SubjectTypeRequest struct {
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}

type SubjectTypeResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}
