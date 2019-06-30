package service

import "awesome-portal-api/internal/storage"

type TypeSubService struct {
	storage.TypeSubStorage
}

func (t *TypeSubService) Save(typeSub *storage.TypeSub) bool {
	return t.TypeSubStorage.Save(typeSub)
}
