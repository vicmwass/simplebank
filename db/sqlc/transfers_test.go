package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"trail.com/simplebank/util"
)

func createRandomTransfer(t *testing.T, from_account Account, to_account Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: from_account.ID,
		ToAccountID:   to_account.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToAccountID, transfer.ToAccountID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer

}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	createRandomTransfer(t, account1, account2)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)
	transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, transfer2)

	require.Equal(t, transfer1.ID, transfer2.ID)
	require.Equal(t, transfer1.FromAccountID, transfer2.FromAccountID)
	require.Equal(t, transfer1.ToAccountID, transfer2.ToAccountID)
	require.Equal(t, transfer1.Amount, transfer2.Amount)

	require.WithinDuration(t, transfer1.CreatedAt, transfer2.CreatedAt, time.Second)
}

func TestGetTransferWithTAIandFAI(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	for i := 0; i < 3; i++ {
		createRandomTransfer(t, account1, account2)
	}

	transfers1, err1 := testQueries.GetTransferByFAI(context.Background(), account1.ID)
	require.NoError(t, err1)
	require.NotEmpty(t, transfers1)
	require.Len(t, transfers1, 3)

	for _, transfer := range transfers1 {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}

	transfers2, err2 := testQueries.GetTransferByTAI(context.Background(), account2.ID)
	require.NoError(t, err2)
	require.NotEmpty(t, transfers2)
	require.Len(t, transfers1, 3)

	for _, transfer := range transfers2 {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}

	arg := GetTransferByTAIandFAIParams{
		FromAccountID: account1.ID,
		ToAccountID:   account2.ID,
	}

	transfers3, err3 := testQueries.GetTransferByTAIandFAI(context.Background(), arg)
	require.NoError(t, err3)
	require.NotEmpty(t, transfers3)
	require.Len(t, transfers1, 3)

	for _, transfer := range transfers3 {
		require.NotEmpty(t, transfer)
		require.Equal(t, transfer.FromAccountID, account1.ID)
		require.Equal(t, transfer.ToAccountID, account2.ID)
	}

}

func TestGetListTransfers(t *testing.T) {
	for i := 0; i < 5; i++ {
		account1 := createRandomAccount(t)
		account2 := createRandomAccount(t)
		createRandomTransfer(t, account1, account2)
	}

	transfers, err := testQueries.ListTransfers(context.Background(), 5)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}

}

func TestDeleteTransfers(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	transfer1 := createRandomTransfer(t, account1, account2)

	err := testQueries.DeleteTransfer(context.Background(), transfer1.ID)
	require.NoError(t, err)

	transfer2, err2 := testQueries.GetTransfer(context.Background(), transfer1.ID)
	require.Error(t, err2)
	require.Empty(t, transfer2)
	require.EqualError(t, err2, sql.ErrNoRows.Error())

}
