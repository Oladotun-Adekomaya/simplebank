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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {

}
