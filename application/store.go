package application

import (
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/domain/repository"
)

func NewStoreService(repo repository.StoreRepository) *storeService {
	return &storeService{
		DBRepository: repo,
	}
}

type storeService struct {
	DBRepository repository.StoreRepository
}

func (s *storeService) Get(storeID domain.StoreID) (domain.Store, error) {
	return s.DBRepository.Get(storeID)
}

func (s *storeService) GetList() ([]domain.Store, error) {
	return s.DBRepository.GetList()
}
