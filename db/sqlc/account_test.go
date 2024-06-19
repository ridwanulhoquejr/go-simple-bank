package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    "Ridwan",
		Balance:  1000,
		Currency: "TK",
	}

	act, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, act)

	require.Equal(t, arg.Owner, act.Owner)
	require.Equal(t, arg.Balance, act.Balance)
	require.Equal(t, arg.Currency, act.Currency)

	require.NotZero(t, act.ID)
	require.NotZero(t, act.CreatedAt)
}
