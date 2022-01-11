package db

import (
	"context"
	"simpleBank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T) Transfer {
    account1Id := getAccountId(t)[0]
    account2Id := getAccountId(t)[1]

    arg := CreateTransferParams{
        FromAccountID: account1Id,
        ToAccountID:   account2Id,
        Amount:        util.RandomMoney(),
    }

    transfer, err := testQueries.CreateTransfer(context.Background(), arg)
    require.NoError(t, err)
    require.NotEmpty(t, transfer)

    require.Equal(t, arg.Amount, transfer.Amount)
    require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
    require.Equal(t, arg.ToAccountID, transfer.ToAccountID)

    require.NotZero(t, transfer.ID)
    require.NotZero(t, transfer.CreatedAt)

    return transfer
}

func TestCreateTransfer(t *testing.T) {
    createRandomTransfer(t)
}

func TestGetTransfer(t *testing.T) {
    transfer1 := createRandomTransfer(t)

    transfer2, err := testQueries.GetTransfer(context.Background(), transfer1.ID)
    require.NoError(t, err)
    require.NotEmpty(t, transfer2)

    require.Equal(t, transfer2.Amount, transfer1.Amount)
    require.Equal(t, transfer2.FromAccountID, transfer1.FromAccountID)
    require.Equal(t, transfer2.ToAccountID, transfer1.ToAccountID)

    require.NotZero(t, transfer2.ID)
    require.WithinDuration(t, transfer2.CreatedAt, transfer1.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
    for i := 0; i < 10; i++ {
        createRandomTransfer(t)
    }

    arg := ListTransfersParams{
        FromAccountID: getAccountId(t)[0],
        ToAccountID:   getAccountId(t)[1],
        Limit:         5,
        Offset:        5,
    }

    transfers, err := testQueries.ListTransfers(context.Background(), arg)
    require.NoError(t, err)
    require.Len(t, transfers, 5)

    for _, transfer := range transfers {
        require.NotEmpty(t, transfer)
    }
}