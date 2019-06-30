package dtos

type SubjectTypeRequest struct {
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}

type SubjectTypeResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}
