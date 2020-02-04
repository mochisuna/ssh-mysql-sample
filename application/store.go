package application

import (
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/domain/repository"
	"github.com/mochisuna/ssh-mysql-sample/domain/service"
)

func NewStoreService(repo repository.StoreRepository) service.StoreService {
	return &storeService{
		DBRepository: repo,
	}
}

type storeService struct {
	DBRepository repository.StoreRepository
}

func (s *storeService) Get(storeID domain.StoreID) (*domain.Store, error) {
	ret, err := s.DBRepository.Get(storeID)
	if err != nil {
		return nil, err
	}
	return &domain.Store{
		ID:        ret.ID,
		UID:       ret.UID,
		Name:      ret.Name,
		Status:    ret.Status,
		CreatedAt: ret.CreatedAt,
		UpdatedAt: ret.UpdatedAt,
	}, nil
}

func (s *storeService) GetList() ([]domain.Store, error) {
	ret, err := s.DBRepository.GetList()
	if err != nil {
		return nil, err
	}
	stores := make([]domain.Store, len(ret))
	for i, store := range ret {
		stores[i] = domain.Store{
			ID:        store.ID,
			UID:       store.UID,
			Name:      store.Name,
			Status:    store.Status,
			CreatedAt: store.CreatedAt,
			UpdatedAt: store.UpdatedAt,
		}
	}
	return stores, nil
}
