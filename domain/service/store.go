package service

import "github.com/mochisuna/ssh-mysql-sample/domain"

type StoreService interface {
	Get(domain.StoreID) (*domain.Store, error)
	GetList() ([]domain.Store, error)
}
