package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)

	// Create two accounts
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)

	// run n concurrent transfer transactions
	n := 5
	amount := int64(10)

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
			errs <- err
			results <- result
		}()
	}

	// check results
	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		// check transfer
		require.NotEmpty(t, result.Transfer)
		require.Equal(t, account1.ID, result.Transfer.FromAccountID)
		require.Equal(t, account2.ID, result.Transfer.ToAccountID)
		require.Equal(t, amount, result.Transfer.Amount)
		require.NotZero(t, result.Transfer.ID)
		require.NotZero(t, result.Transfer.CreatedAt)

		_, err = store.GetTransfer(context.Background(), result.Transfer.ID)

		// check entries
		require.NotEmpty(t, result.FromEntry)
		require.NotEmpty(t, result.ToEntry)
		require.Equal(t, account1.ID, result.FromEntry.AccountID)
		require.Equal(t, account2.ID, result.ToEntry.AccountID)
		require.Equal(t, -amount, result.FromEntry.Amount)
		require.Equal(t, amount, result.ToEntry.Amount)
		require.NotZero(t, result.FromEntry.ID)
		require.NotZero(t, result.FromEntry.CreatedAt)

		_, err = store.GetEntry(context.Background(), result.FromEntry.ID)

		// check accounts
		// fromAccount, err := store.GetAccount(context.Background(), account1.ID)
		// require.NoError(t, err)
		// require.NotEmpty(t, fromAccount)
		// require.Equal(t, account1.Balance-amount, fromAccount.Balance)

		// toAccount, err := store.GetAccount(context.Background(), account2.ID)
		// require.NoError(t, err)
		// require.NotEmpty(t, toAccount)
		// require.Equal(t, account2.Balance+amount, toAccount.Balance)
	}
}
