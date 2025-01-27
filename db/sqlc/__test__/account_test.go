package test

import (
	"context"
	"testing"

	db "github.com/morshedulmunna/simple_bank/db/sqlc"
)

func TestCreateAccount(t *testing.T) {
	arg := db.CreateAccountParams{
		Owner:    "test",
		Balance:  1000,
		Currency: "USD",
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	if account.Owner != arg.Owner {
		t.Errorf("owner: got %v want %v", account.Owner, arg.Owner)
	}
	if account.Balance != arg.Balance {
		t.Errorf("balance: got %v want %v", account.Balance, arg.Balance)
	}
	if account.Currency != arg.Currency {
		t.Errorf("currency: got %v want %v", account.Currency, arg.Currency)
	}
	if account.ID == 0 {
		t.Error("id should not be 0")
	}
	if account.CreatedAt.Time.IsZero() {
		t.Error("created_at should not be zero")
	}
}
