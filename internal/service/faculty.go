package service

import "awesome-portal-api/internal/storage"

type FacultyService struct {
	storage.FacultyStorage
}

func (f *FacultyService) FacultyByShort(short string) (*storage.Faculty, bool) {
	return f.FacultyStorage.FacultyByShort(short)
}

func (f *FacultyService) Save(faculty *storage.Faculty) bool {
	return f.FacultyStorage.Save(faculty)
}
