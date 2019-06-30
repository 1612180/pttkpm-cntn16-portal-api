package storage

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Score struct {
	ID       int     `json:"id"`
	Midterm  float64 `json:"midterm"`
	Endterm  float64 `json:"endterm"`
	Final    float64 `json:"final"`
	EnrollID int     `json:"enroll_id"`
}

type ScoreStorage interface {
	ScoreByEnrollID(enrollID int) (*Score, bool)
	Save(score *Score) bool
}

type ScoreGorm struct {
	*gorm.DB
}

func NewScoreStorage(db *gorm.DB) ScoreStorage {
	return &ScoreGorm{DB: db}
}

func (s *ScoreGorm) ScoreByEnrollID(enrollID int) (*Score, bool) {
	var score Score
	if err := s.DB.Where("enroll_id = ?", enrollID).First(&score).Error; err != nil {
		log.Println(err)
		return nil, false
	}
	return &score, true
}

func (s *ScoreGorm) Save(score *Score) bool {
	var oldScore Score
	if err := s.DB.Where("enroll_id = ?", score.EnrollID).First(&oldScore).Error; err != nil {
		log.Println(err)
		return false
	}

	oldScore.Midterm = score.Midterm
	oldScore.Endterm = score.Endterm
	oldScore.Final = score.Final
	if err := s.DB.Save(&oldScore).Error; err != nil {
		log.Println(err)
		return false
	}
	return true
}
