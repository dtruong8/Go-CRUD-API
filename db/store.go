package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	*CustomerStore
}

func NewStore(dataSourceName string) (*Store, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Store{
		CustomerStore: &CustomerStore{DB: db},
	}, nil
}
