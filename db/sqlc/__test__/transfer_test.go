package test

import (
	"context"
	"testing"

	db "github.com/morshedulmunna/simple_bank/db/sqlc"
	"github.com/morshedulmunna/simple_bank/utils"
)

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	arg := db.CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if transfer.FromAccountID != arg.FromAccountID {
		t.Errorf("from_account_id: got %v want %v", transfer.FromAccountID, arg.FromAccountID)
	}
	if transfer.ToAccountID != arg.ToAccountID {
		t.Errorf("to_account_id: got %v want %v", transfer.ToAccountID, arg.ToAccountID)
	}
	if transfer.Amount != arg.Amount {
		t.Errorf("amount: got %v want %v", transfer.Amount, arg.Amount)
	}
	if transfer.ID == 0 {
		t.Error("id should not be 0")
	}
	if transfer.CreatedAt.Time.IsZero() {
		t.Error("created_at should not be zero")
	}
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
	if err != nil {
		t.Fatal(err)
	}

	if transfer1.ID != transfer2.ID {
		t.Errorf("id: got %v want %v", transfer2.ID, transfer1.ID)
	}
	if transfer1.FromAccountID != transfer2.FromAccountID {
		t.Errorf("from_account_id: got %v want %v", transfer2.FromAccountID, transfer1.FromAccountID)
	}
	if transfer1.ToAccountID != transfer2.ToAccountID {
		t.Errorf("to_account_id: got %v want %v", transfer2.ToAccountID, transfer1.ToAccountID)
	}
	if transfer1.Amount != transfer2.Amount {
		t.Errorf("amount: got %v want %v", transfer2.Amount, transfer1.Amount)
	}
}

func TestListTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, account1, account2)
	}

	arg := db.ListTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfer(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if len(transfers) != 5 {
		t.Errorf("expected 5 transfers, got %d", len(transfers))
	}

	for _, transfer := range transfers {
		if transfer.ID == 0 {
			t.Error("transfer ID should not be 0")
		}
		if transfer.FromAccountID != account1.ID {
			t.Errorf("expected from account ID %v, got %v", account1.ID, transfer.FromAccountID)
		}
		if transfer.ToAccountID != account2.ID {
			t.Errorf("expected to account ID %v, got %v", account2.ID, transfer.ToAccountID)
		}
	}
}

// Helper function to create a random transfer for testing
func createRandomTransfer(t *testing.T, account1, account2 db.Account) db.Transfer {
	arg := db.CreateTransferParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
		Amount:        utils.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	return transfer
}
