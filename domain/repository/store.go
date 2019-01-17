package repository

import "github.com/mochisuna/ssh-mysql-sample/domain"

type StoreRepository interface {
	Get(domain.StoreID) (domain.Store, error)
	GetList() ([]domain.Store, error)
}
