package dtos

type ProgramRequest struct {
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}

type ProgramResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short_name"`
	LongName  string `json:"long_name"`
}
