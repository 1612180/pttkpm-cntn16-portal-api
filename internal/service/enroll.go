package service

import "awesome-portal-api/internal/storage"

type EnrollService struct {
	storage.EnrollStorage
}

type MultiEnroll struct {
	StudentID  int   `json:"student_id"`
	SubjectIDS []int `json:"subject_ids"`
}

func (e *EnrollService) Save(multiEnroll *MultiEnroll) bool {
	status := false
	for _, subjectID := range multiEnroll.SubjectIDS {
		enroll := storage.Enroll{
			StudentID: multiEnroll.StudentID,
			SubjectID: subjectID,
		}
		// one true is enough
		if ok := e.EnrollStorage.Save(&enroll); ok {
			status = true
		}
	}
	return status
}
