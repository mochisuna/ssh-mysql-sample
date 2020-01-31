package domain

import (
	"time"
)

type StoreID int

const StoreTableName = "stores"

type Store struct {
	ID        StoreID
	UID       string
	Name      string
	Status    int
	CreatedAt time.Time
	UpdatedAt time.Time
}
