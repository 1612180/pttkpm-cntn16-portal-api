package repositories

import (
	"awesome-portal-api/internal/models"
	"os"

	"github.com/jinzhu/gorm"
)

type StudentRepo interface {
	FetchAll() ([]*models.Student, error)
	FindByID(id int) (*models.Student, error)
	FindByMSSV(mssv string) (*models.Student, error)
	Create(student *models.Student, account *models.Account) error
	DeleteByMSSV(mssv string) error
}

type AccountRepo interface {
	FindByID(id int) (*models.Account, error)
}

type SubjectRepo interface {
	FetchAll() ([]*models.Subject, error)
	FindByID(id int) (*models.Subject, error)
	Create(subject *models.Subject) error
	DeleteByID(id int) error
}

type SubjectTypeRepo interface {
	FetchAll() ([]*models.SubjectType, error)
	FindByID(id int) (*models.SubjectType, error)
	FindByShort(short string) (*models.SubjectType, error)
	Create(subjectType models.SubjectType) error
	DeleteByID(id int) error
	DeleteByShort(short string) error
}

type SubjectPrerequisiteRepo interface {
	FetchAll() ([]*models.SubjectPrerequisite, error)
	FindByID(id int) (*models.SubjectPrerequisite, error)
	Create(subjectType models.SubjectPrerequisite) error
	DeleteByID(id int) error
}

type ProgramRepo interface {
	FetchAll() ([]*models.Program, error)
	FindByID(id int) (*models.Program, error)
	FindByShort(short string) (*models.Program, error)
	Create(program *models.Program) error
	DeleteByID(id int) error
	DeleteByShort(short string) error
}

type FacultyRepo interface {
	FetchAll() ([]*models.Faculty, error)
	FindByID(id int) (*models.Faculty, error)
	FindByShort(short string) (*models.Faculty, error)
	Create(program *models.Faculty) error
	DeleteByID(id int) error
	DeleteByShort(short string) error
}

type Repos interface {
	CreateAll() (StudentRepo, AccountRepo, ProgramRepo, FacultyRepo)
}

type ReposGorm struct {
	*gorm.DB
}

func NewReposGorm(db *gorm.DB) Repos {
	return &ReposGorm{DB: db}
}

func (r *ReposGorm) CreateAll() (StudentRepo, AccountRepo, ProgramRepo, FacultyRepo) {
	if os.Getenv("DATABASE_MODE") == "debug" {
		r.DB.DropTableIfExists(&models.Account{})
		r.DB.DropTableIfExists(&models.Student{})
		r.DB.DropTableIfExists(&models.Program{})
		r.DB.DropTableIfExists(&models.Faculty{})

		r.DB.DropTableIfExists(&models.Subject{})
		r.DB.DropTableIfExists(&models.SubjectType{})
		r.DB.DropTableIfExists(&models.SubjectPrerequisite{})
		r.DB.DropTableIfExists(&models.StudentSubject{})
	}

	r.DB.AutoMigrate(&models.Account{})
	r.DB.AutoMigrate(&models.Student{})
	r.DB.AutoMigrate(&models.Program{})
	r.DB.AutoMigrate(&models.Faculty{})

	r.DB.AutoMigrate(&models.Subject{})
	r.DB.AutoMigrate(&models.SubjectType{})
	r.DB.AutoMigrate(&models.SubjectPrerequisite{})
	r.DB.AutoMigrate(&models.StudentSubject{})

	return &StudentGorm{DB: r.DB},
		&AccountGorm{DB: r.DB},
		&ProgramGorm{DB: r.DB},
		&FacultyGorm{DB: r.DB}
}
