package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/ridwanulhoquejr/go-simple-bank/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomBalance(),
		Currency: util.RandomCurrency(),
	}

	act, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, act)

	require.Equal(t, arg.Owner, act.Owner)
	require.Equal(t, arg.Balance, act.Balance)
	require.Equal(t, arg.Currency, act.Currency)

	require.NotZero(t, act.ID)
	require.NotZero(t, act.CreatedAt)

	return act
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

// test for get account
func TestGetAccoun(t *testing.T) {
	act1 := CreateRandomAccount(t)

	act2, err := testQueries.GetAccount(context.Background(), act1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, act2)

	require.Equal(t, act1.ID, act2.ID)
	require.Equal(t, act1.Owner, act2.Owner)
	require.Equal(t, act1.Balance, act2.Balance)
	require.Equal(t, act1.Currency, act2.Currency)
	require.WithinDuration(t, act1.CreatedAt, act2.CreatedAt, time.Second)
}

// test for update account
func TestUpdateAccount(t *testing.T) {
	act1 := CreateRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      act1.ID,
		Balance: util.RandomBalance(),
	}

	act2, err := testQueries.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, act2)

	require.Equal(t, act1.ID, act2.ID)
	require.Equal(t, act1.Owner, act2.Owner)

	// the balance should not be equal for act1 and act2;
	// so we compare with  arg.Balance
	require.Equal(t, arg.Balance, act2.Balance)
	require.Equal(t, act1.Currency, act2.Currency)
	require.WithinDuration(t, act1.CreatedAt, act2.CreatedAt, time.Second)
}

// test for delete account
func TestDeleteAccount(t *testing.T) {
	act1 := CreateRandomAccount(t)

	err := testQueries.DeleteAccount(context.Background(), act1.ID)
	require.NoError(t, err)

	act2, err := testQueries.GetAccount(context.Background(), act1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, act2)
}

// test list accounts
func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	acts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, acts, 5)

	for _, act := range acts {
		require.NotEmpty(t, act)
	}
}
