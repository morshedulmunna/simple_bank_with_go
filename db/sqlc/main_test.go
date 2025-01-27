package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries

const (
	dbSource = "postgresql://root:secrets@localhost:5432/simple_bank_db?sslmode=disable"
)

func TestMain(m *testing.M) {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}
	defer conn.Close()

	testQueries = New(conn)
	os.Exit(m.Run())
}
