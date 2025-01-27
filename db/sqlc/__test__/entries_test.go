package test

import (
	"context"
	"testing"

	db "github.com/morshedulmunna/simple_bank/db/sqlc"
	"github.com/morshedulmunna/simple_bank/utils"
)

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	arg := db.CreateEntriesParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if entry.AccountID != arg.AccountID {
		t.Errorf("account_id: got %v want %v", entry.AccountID, arg.AccountID)
	}
	if entry.Amount != arg.Amount {
		t.Errorf("amount: got %v want %v", entry.Amount, arg.Amount)
	}
	if entry.ID == 0 {
		t.Error("id should not be 0")
	}
	if entry.CreatedAt.Time.IsZero() {
		t.Error("created_at should not be zero")
	}
}

func TestGetEntry(t *testing.T) {
	// Create entry first
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)

	if err != nil {
		t.Fatal(err)
	}

	if entry1.ID != entry2.ID {
		t.Errorf("id: got %v want %v", entry2.ID, entry1.ID)
	}
	if entry1.AccountID != entry2.AccountID {
		t.Errorf("account_id: got %v want %v", entry2.AccountID, entry1.AccountID)
	}
	if entry1.Amount != entry2.Amount {
		t.Errorf("amount: got %v want %v", entry2.Amount, entry1.Amount)
	}
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		createRandomEntry(t, account)
	}

	entries, err := testQueries.ListEntries(context.Background(), db.ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(entries) != 5 {
		t.Errorf("expected 5 entries, got %d", len(entries))
	}

	for _, entry := range entries {
		if entry.ID == 0 {
			t.Error("entry ID should not be 0")
		}
		if entry.AccountID != account.ID {
			t.Errorf("expected account ID %v, got %v", account.ID, entry.AccountID)
		}
	}
}

// Helper function to create a random entry for testing
func createRandomEntry(t *testing.T, account db.Account) db.Entry {
	arg := db.CreateEntriesParams{
		AccountID: account.ID,
		Amount:    utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntries(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	return entry
}
