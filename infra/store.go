package infra

import (
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mochisuna/ssh-mysql-sample/domain"
	"github.com/mochisuna/ssh-mysql-sample/infra/mysql"
)

const (
	// Table
	storeTable = "stores"
)

type storeColumns struct {
	ID        domain.StoreID `db:"id"`
	UID       string         `db:"uid"`
	Name      string         `db:"name"`
	Status    int            `db:"status"`
	CreatedAt time.Time      `db:"created_at"`
	UpdatedAt time.Time      `db:"updated_at"`
}

type storeRepository struct {
	DBClient *mysql.Client
}

func NewStoreRepository(dbClient *mysql.Client) *storeRepository {
	return &storeRepository{
		DBClient: dbClient,
	}
}

func (r *storeRepository) Get(storeID domain.StoreID) (domain.Store, error) {
	columns := []string{
		"id",
		"uid",
		"name",
		"status",
		"created_at",
		"updated_at",
	}
	ret := domain.Store{}
	var store storeColumns
	err := sq.Select(columns...).
		From("stores").
		Where(sq.Eq{
			"stores.id": storeID,
		}).
		RunWith(r.DBClient.DB).
		QueryRow().
		Scan(
			&store.ID,
			&store.UID,
			&store.Name,
			&store.Status,
			&store.CreatedAt,
			&store.UpdatedAt,
		)
	if err != nil {
		return ret, err
	}
	ret = domain.Store{
		ID:        store.ID,
		UID:       store.UID,
		Name:      store.Name,
		Status:    store.Status,
		CreatedAt: store.CreatedAt,
		UpdatedAt: store.UpdatedAt,
	}
	return ret, err
}

func (r *storeRepository) GetList() ([]domain.Store, error) {
	columns := []string{
		"id",
		"uid",
		"name",
		"status",
		"created_at",
		"updated_at",
	}
	ret := []domain.Store{}
	rows, err := sq.Select(columns...).
		From("stores").
		RunWith(r.DBClient.DB).
		Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var store storeColumns
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
		ret = append(ret, domain.Store{
			ID:        store.ID,
			UID:       store.UID,
			Name:      store.Name,
			Status:    store.Status,
			CreatedAt: store.CreatedAt,
			UpdatedAt: store.UpdatedAt,
		})
	}
	return ret, err
}
