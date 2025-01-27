package test

import (
	"context"
	"testing"

	db "github.com/morshedulmunna/simple_bank/db/sqlc"
	"github.com/morshedulmunna/simple_bank/utils"
)

func TestCreateAccount(t *testing.T) {
	arg := db.CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
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

func TestGetAccount(t *testing.T) {
	// Create account first
	account1 := createRandomAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	if err != nil {
		t.Fatal(err)
	}

	if account1.ID != account2.ID {
		t.Errorf("id: got %v want %v", account2.ID, account1.ID)
	}
	if account1.Owner != account2.Owner {
		t.Errorf("owner: got %v want %v", account2.Owner, account1.Owner)
	}
	if account1.Balance != account2.Balance {
		t.Errorf("balance: got %v want %v", account2.Balance, account1.Balance)
	}
	if account1.Currency != account2.Currency {
		t.Errorf("currency: got %v want %v", account2.Currency, account1.Currency)
	}
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	deletedAccount, err := testQueries.DeleteAccount(context.Background(), account1.ID)

	if err != nil {
		t.Fatal(err)
	}

	if account1.ID != deletedAccount.ID {
		t.Errorf("id: got %v want %v", deletedAccount.ID, account1.ID)
	}

	// Verify account is deleted by trying to get it
	_, err = testQueries.GetAccount(context.Background(), account1.ID)
	if err == nil {
		t.Error("expected error getting deleted account")
	}
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	accounts, err := testQueries.ListAccounts(context.Background(), db.ListAccountsParams{
		Limit:  5,
		Offset: 0,
	})
	if err != nil {
		t.Fatal(err)
	}

	if len(accounts) != 5 {
		t.Errorf("expected 5 accounts, got %d", len(accounts))
	}

	for _, account := range accounts {
		if account.ID == 0 {
			t.Error("account ID should not be 0")
		}
	}
}

// Helper function to create a random account for testing
func createRandomAccount(t *testing.T) db.Account {
	arg := db.CreateAccountParams{
		Owner:    utils.RandomOwner(),
		Balance:  utils.RandomMoney(),
		Currency: utils.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	if err != nil {
		t.Fatal(err)
	}

	return account
}
