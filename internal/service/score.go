package service

import "awesome-portal-api/internal/storage"

type ScoreService struct {
	storage.ScoreStorage
}

func (s *ScoreService) Save(score *storage.Score) bool {
	return s.ScoreStorage.Save(score)
}
