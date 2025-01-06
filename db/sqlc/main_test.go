package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
)

const DATABASE_URL = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries

func TestMain(m *testing.M) {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		log.Fatal("Unable to connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
	// defer conn.Close(context.Background())

}
