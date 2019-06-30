package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type TypeSub struct {
	ID    int    `json:"id"`
	Short string `json:"short"`
	Long  string `json:"long"`
}

type TypeSubStorage interface {
	TypeSub(id int) (*TypeSub, bool)
	TypeSubByShort(short string) (*TypeSub, bool)
	Save(typeSub *TypeSub) bool
}

type TypeSubGorm struct {
	*gorm.DB
}

func NewTypeSubStorage(db *gorm.DB) TypeSubStorage {
	return &TypeSubGorm{DB: db}
}

func (t *TypeSubGorm) TypeSub(id int) (*TypeSub, bool) {
	var typeSub TypeSub
	if err := t.DB.Where("id = ?", id).First(&typeSub).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &typeSub, true
}

func (t *TypeSubGorm) TypeSubByShort(short string) (*TypeSub, bool) {
	var typeSub TypeSub
	if err := t.DB.Where("short = ?", short).First(&typeSub).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &typeSub, true
}

func (t *TypeSubGorm) Save(typeSub *TypeSub) bool {
	if err := t.DB.Create(typeSub).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
