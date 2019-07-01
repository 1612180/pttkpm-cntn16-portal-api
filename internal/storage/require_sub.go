package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type RequireSub struct {
	ID           int `json:"id"`
	SubjectPreID int `json:"subject_pre_id"`
	SubjectCurID int `json:"subject_cur_id"`
}

type RequireSubStorage interface {
	RequireSubByCurID(subjectCurID int) ([]*RequireSub, bool)
	Save(requireSub *RequireSub) bool
}

type RequireSubGorm struct {
	*gorm.DB
}

func NewRequireSubStorage(db *gorm.DB) RequireSubStorage {
	return &RequireSubGorm{DB: db}
}

func (r *RequireSubGorm) RequireSubByCurID(subjectCurID int) ([]*RequireSub, bool) {
	var requireSubs []*RequireSub
	if err := r.DB.Where("subject_cur_id = ?", subjectCurID).Find(&requireSubs).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return requireSubs, true
}

func (r *RequireSubGorm) Save(requireSub *RequireSub) bool {
	if err := r.DB.Create(requireSub).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
