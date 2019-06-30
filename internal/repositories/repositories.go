package repositories

import (
	"awesome-portal-api/internal/models"
	"os"

	"github.com/jinzhu/gorm"
)

type StudentRepo interface {
	FetchAll() ([]*models.Student, bool)
	FindByID(id int) (*models.Student, bool)
	FindByMSSV(mssv string) (*models.Student, bool)
	Create(student *models.Student, account *models.Account) bool
	DeleteByMSSV(mssv string) bool
}

type AccountRepo interface {
	FindByID(id int) (*models.Account, bool)
}

type SubjectRepo interface {
	FetchAll() ([]*models.Subject, bool)
	FindByID(id int) (*models.Subject, bool)
	Create(subject *models.Subject) bool
	DeleteByID(id int) bool
}

type SubjectTypeRepo interface {
	FetchAll() ([]*models.SubjectType, bool)
	FindByID(id int) (*models.SubjectType, bool)
	FindByShort(short string) (*models.SubjectType, bool)
	Create(subjectType *models.SubjectType) bool
	DeleteByID(id int) bool
	DeleteByShort(short string) bool
}

type SubjectPrerequisiteRepo interface {
	FetchAll() ([]*models.SubjectPrerequisite, bool)
	FindByID(id int) (*models.SubjectPrerequisite, bool)
	Create(subjectType *models.SubjectPrerequisite) bool
	DeleteByID(id int) bool
}

type ProgramRepo interface {
	FetchAll() ([]*models.Program, bool)
	FindByID(id int) (*models.Program, bool)
	FindByShort(short string) (*models.Program, bool)
	Create(program *models.Program) bool
	DeleteByID(id int) bool
	DeleteByShort(short string) bool
}

type FacultyRepo interface {
	FetchAll() ([]*models.Faculty, bool)
	FindByID(id int) (*models.Faculty, bool)
	FindByShort(short string) (*models.Faculty, bool)
	Create(program *models.Faculty) bool
	DeleteByID(id int) bool
	DeleteByShort(short string) bool
}

type Repos interface {
	CreateAll() (
		StudentRepo,
		AccountRepo,
		ProgramRepo,
		FacultyRepo,
		SubjectRepo,
		SubjectTypeRepo,
	)
}

type ReposGorm struct {
	*gorm.DB
}

func NewReposGorm(db *gorm.DB) Repos {
	return &ReposGorm{DB: db}
}

func (r *ReposGorm) CreateAll() (
	StudentRepo,
	AccountRepo,
	ProgramRepo,
	FacultyRepo,
	SubjectRepo,
	SubjectTypeRepo,
) {
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
		&FacultyGorm{DB: r.DB},
		&SubjectGorm{DB: r.DB},
		&SubjectTypeGorm{DB: r.DB}
}
