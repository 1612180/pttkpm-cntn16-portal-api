package repositories

import (
	"awesome-portal-api/internal/models"

	"github.com/jinzhu/gorm"
)

type StudentRepo interface {
	FetchAll() ([]*models.Student, error)
	FindByID(id int) (*models.Student, error)
	FindByMSSV(mssv string) (*models.Student, error)
	Create(student *models.Student, account *models.Account) error
	// UpdateInfo(student *models.Student) error
	// UpdatePassword(account *models.Account) error
	Delete(mssv string) error
}

type AccountRepo interface {
	FindByID(id int) (*models.Account, error)
}

type SubjectRepo interface {
	FetchAll() ([]*models.Subject, error)
	FindByID(id int) (*models.Subject, error)
	Create(subject *models.Subject) error
	Delete(id int) error
}

type SubjectTypeRepo interface {
	FetchAll() ([]*models.SubjectType, error)
	FindByID(id int) *models.SubjectType
}

type Repos interface {
	CreateAll() (StudentRepo, AccountRepo)
}

type ReposGorm struct {
	*gorm.DB
}

func NewReposGorm(db *gorm.DB) Repos {
	return &ReposGorm{DB: db}
}

func (r *ReposGorm) CreateAll() (StudentRepo, AccountRepo) {
	r.DB.DropTableIfExists(&models.Account{})
	r.DB.DropTableIfExists(&models.Student{})
	r.DB.DropTableIfExists(&models.Program{})
	r.DB.DropTableIfExists(&models.Faculty{})

	r.DB.AutoMigrate(&models.Account{})
	r.DB.AutoMigrate(&models.Student{})
	r.DB.AutoMigrate(&models.Program{})
	r.DB.AutoMigrate(&models.Faculty{})

	return &StudentGorm{DB: r.DB}, &AccountGorm{DB: r.DB}
}
