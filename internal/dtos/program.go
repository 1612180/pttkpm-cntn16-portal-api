package dtos

type ProgramRequest struct {
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}

type ProgramResponse struct {
	ID        int    `json:"id"`
	ShortName string `json:"short"`
	LongName  string `json:"long"`
}
