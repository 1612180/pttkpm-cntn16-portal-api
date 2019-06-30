package service

import "awesome-portal-api/internal/storage"

type ProgramService struct {
	storage.ProgramStorage
}

func (p *ProgramService) ProgramByShort(short string) (*storage.Program, bool) {
	return p.ProgramStorage.ProgramByShort(short)
}

func (p *ProgramService) Save(program *storage.Program) bool {
	return p.ProgramStorage.Save(program)
}
