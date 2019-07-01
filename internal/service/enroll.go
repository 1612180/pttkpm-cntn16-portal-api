package service

import "awesome-portal-api/internal/storage"

type EnrollService struct {
	storage.EnrollStorage
}

type MultiEnroll struct {
	StudentID  int   `json:"student_id"`
	SubjectIDS []int `json:"subject_ids"`
}

func (e *EnrollService) SaveMulti(multiEnroll *MultiEnroll) bool {
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

func (e *EnrollService) SaveTryMulti(multiEnroll *MultiEnroll) bool {
	status := false
	for _, subjectID := range multiEnroll.SubjectIDS {
		tryEnroll := storage.TryEnroll{
			StudentID: multiEnroll.StudentID,
			SubjectID: subjectID,
		}
		// one true only
		if ok := e.EnrollStorage.SaveTry(&tryEnroll); ok {
			status = true
		}
	}
	return status
}

func (e *EnrollService) SaveRealAll() bool {
	tryEnrolls, ok := e.EnrollStorage.TryEnrolls()
	if !ok {
		return false
	}

	status := false
	for _, tryEnroll := range tryEnrolls {
		// one true yeah
		if ok := e.EnrollStorage.SaveReal(tryEnroll); ok {
			status = true
		}
	}
	return status
}

func (e *EnrollService) DeleteTryMulti(multiEnroll *MultiEnroll) bool {
	status := false
	for _, subjectID := range multiEnroll.SubjectIDS {
		// one true only
		if ok := e.EnrollStorage.DeleteTrySSID(multiEnroll.StudentID, subjectID); ok {
			status = true
		}
	}
	return status

}
