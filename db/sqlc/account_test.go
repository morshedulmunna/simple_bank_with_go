package db

import (
	"context"
	"testing"

	"github.com/morshedulmunna/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    utils.RandomString(12),
		Balance:  utils.RandomInt(1, 1000),
		Currency: utils.RandomCurrency(),
	}

	// Test the CreateAccount function here
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotNil(t, account)

	// Add more tests for other functions here...
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)
	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)
	require.NotZero(t, account.UpdatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	deleteAccount, err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotNil(t, deleteAccount)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotNil(t, account2)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: utils.RandomInt(1, 1000),
	}

	err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)

}

func TestListAccount(t *testing.T) {
	createRandomAccount(t)
	arg := ListAccountsParams{
		Limit:  1,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
	require.Len(t, accounts, 1)
}
