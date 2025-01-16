package db

import "github.com/jackc/pgx/v5"

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *pgx.Conn
}
