package infra

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/domain/repository"
	"github.com/mochisuna/ssh-mysql-sample/infra/mysql"
)

const (
	// Table
	storeTable = "stores"
)

type storeRepository struct {
	DBClient *mysql.Client
}

func NewStoreRepository(dbClient *mysql.Client) *storeRepository {
	return &storeRepository{
		DBClient: dbClient,
	}
}

func (r *storeRepository) Get(storeID domain.StoreID) (repository.Store, error) {
	columns := []string{
		"id",
		"uid",
		"name",
		"status",
		"created_at",
		"updated_at",
	}
	ret := repository.Store{}
	err := sq.Select(columns...).
		From("stores").
		Where(sq.Eq{
			"stores.id": storeID,
		}).
		RunWith(r.DBClient.DB).
		QueryRow().
		Scan(
			&ret.ID,
			&ret.UID,
			&ret.Name,
			&ret.Status,
			&ret.CreatedAt,
			&ret.UpdatedAt,
		)
	if err != nil {
		return ret, err
	}

	return ret, err
}

func (r *storeRepository) GetList() ([]repository.Store, error) {
	columns := []string{
		"id",
		"uid",
		"name",
		"status",
		"created_at",
		"updated_at",
	}
	ret := []repository.Store{}
	rows, err := sq.Select(columns...).
		From("stores").
		RunWith(r.DBClient.DB).
		Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var store repository.Store
		err = rows.Scan(
			&store.ID,
			&store.UID,
			&store.Name,
			&store.Status,
			&store.CreatedAt,
			&store.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		ret = append(ret, store)
	}
	return ret, err
}
