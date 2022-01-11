package db

import (
	"context"
	"math/rand"
	"simpleBank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func getRandomId(t *testing.T) int64 {
	arg := ListAccountsParams{
		Limit: 10,
		Offset: 5,
	}
	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	
	var accountId []int64
	for _, account := range accounts {
		accountId = append(accountId, account.ID)
	}
	n := len(accountId)
	return accountId[rand.Intn(n)]
}

func createRandomEntry(t *testing.T) Entry {
    arg := CreateEntryParams{
        AccountID: getRandomId(t),
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
		AccountID: getRandomId(t),
		Limit: 5,
		Offset: 0,
	}

	entries, err := testQueries.ListEntries(context.Background(),  arg)
	require.NoError(t, err)
	require.Len(t, entries, 1)

	for _,entry := range entries {
		require.NotEmpty(t, entry) 
	}
}