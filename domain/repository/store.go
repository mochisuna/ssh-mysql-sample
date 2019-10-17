package repository

import (
	"time"

	"github.com/mochisuna/ssh-mysql-sample/domain"
)

type StoreRepository interface {
	Get(domain.StoreID) (Store, error)
	GetList() ([]Store, error)
}

type Store struct {
	ID        domain.StoreID `db:"id"`
	UID       string         `db:"uid"`
	Name      string         `db:"name"`
	Status    int            `db:"status"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}
