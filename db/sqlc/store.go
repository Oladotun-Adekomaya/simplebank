package db

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *pgx.Conn
}

// NewStore creates a new Store
func NewStore(db *pgx.Conn) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {

		}
	}
}
