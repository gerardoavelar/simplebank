package db

import (
	"context"
	"database/sql"
)

// Store has all functions to execute queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Create a new Store object
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTX executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	return nil //placeholder
}
