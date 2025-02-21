package db

import (
	"context"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	pool, err := pgxpool.New(context.Background(), "postgresql://root:secret@localhost:5432/simple_bank_db?sslmode=disable")
	if err != nil {
		t.Fatalf("Unable to create connection pool: %v", err)
	}
	defer pool.Close()

	store := NewStore(pool)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurrent transfer transaction

	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for range n {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Account:       amount,
			})

			errs <- err
			results <- result

		}()
	}

	// check result
	for range n {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer
		transfer := result.Transfer

		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, account2.ID, transfer.ToAccountID)
	}
}
