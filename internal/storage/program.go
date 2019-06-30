package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Program struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type ProgramStorage interface {
	Program(id int) (*Program, bool)
	ProgramByShort(short string) (*Program, bool)
	Save(program *Program) bool
}

type ProgramGorm struct {
	*gorm.DB
}

func NewProgramStorage(db *gorm.DB) ProgramStorage {
	return &ProgramGorm{DB: db}
}

func (p *ProgramGorm) Program(id int) (*Program, bool) {
	var program Program
	if err := p.DB.Where("id = ?", id).First(&program).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &program, true
}

func (p *ProgramGorm) ProgramByShort(short string) (*Program, bool) {
	var program Program
	if err := p.DB.Where("short = ?", short).First(&program).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &program, true
}

func (p *ProgramGorm) Save(program *Program) bool {
	if err := p.DB.Create(program).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
