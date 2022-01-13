package db

import (
	"context"
	"simpleBank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getAccountId(t *testing.T) []int64 {
	arg := ListAccountsParams{
		Limit: 2,
		Offset: 0,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	
	var accountId []int64
	for _, account := range accounts {
		accountId = append(accountId, account.ID)
	}
	
	return accountId
}

func createRandomEntry(t *testing.T) Entry {
    arg := CreateEntryParams{
        AccountID: getAccountId(t)[0],
        Amount:    util.RandomMoney(),
    }

    entry, err := testQueries.CreateEntry(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, entry)

    require.Equal(t, arg.AccountID, entry.AccountID)
    require.Equal(t, arg.Amount, entry.Amount)

    require.NotZero(t, entry.ID)
    require.NotZero(t, entry.CreatedAt)

    return entry
}
func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID) 
	require.NoError(t, err)
	require.NotEmpty(t, entry1)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {

	arg := ListEntriesParams{
		AccountID: getAccountId(t)[0],
		Limit: 5,
		Offset: 0,
	}

	entries, err := testQueries.ListEntries(context.Background(),  arg)
	require.NoError(t, err)
	require.Len(t, entries, 2)

	for _,entry := range entries {
		require.NotEmpty(t, entry) 
	}
}