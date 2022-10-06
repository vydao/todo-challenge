package db

import "database/sql"

type Store interface {
	Querier
	// TODO: add transaction methods
}

// SQLStore provides all functions to execute SQL queries
type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
