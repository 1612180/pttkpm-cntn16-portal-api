package dtos

type StudentSubjectRequest struct {
	ScoreMidterm float64 `json:"score_midterm"`
	ScoreFinal   float64 `json:"score_final"`
	ScoreRatio   int     `json:"score_ratio"`
	StudentID    int     `json:"student_id"`
	SubjectID    int     `json:"subject_id"`
}

type StudentSubjectResponse struct {
	ID           int     `json:"id"`
	ScoreMidterm float64 `json:"score_midterm"`
	ScoreFinal   float64 `json:"score_final"`
	ScoreRatio   int     `json:"score_ratio"`
	StudentID    int     `json:"student_id"`
	SubjectID    int     `json:"subject_id"`
}
