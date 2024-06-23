package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	arg := CreateTransferParams{
		FromAccountID: 1,
		ToAccountID:   2,
		Amount:        100,
	}

	act, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, act)

	require.Equal(t, arg.FromAccountID, act.FromAccountID)
	require.Equal(t, arg.ToAccountID, act.ToAccountID)
	require.Equal(t, arg.Amount, act.Amount)

	require.NotZero(t, act.ID)
	require.NotZero(t, act.CreatedAt)

}
