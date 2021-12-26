package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Store struct {
	db *sqlx.DB
}

func Connect(source, driver string) (*Store, error) {
	// Connect to database.
	db, err := sqlx.Connect(source, driver)
	if err != nil {
		return nil, err
	}

	store := &Store{
		db: db,
	}

	log.Println("Connected to database!")
	return store, err
}

func (store *Store) PingDB() error {
	return store.db.Ping()
}
