package db

import (
	"context"
	"testing"

	"github.com/morshedulmunna/simple_bank/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
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

	t.Log("CreateAccount test passed")
}
