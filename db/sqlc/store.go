package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all the functions to execute db queries and transactions.
type Store struct {
	*Queries 
	db *sql.DB 
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
		Queries: New(db),
	}
}

// execTx excutes a function within a database transaction 
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil) 
	if err != nil {
		return err 
	}

	q := New(tx) 
	err = fn(q) 
	if err != nil{
		if rbErr := tx.Rollback(); rbErr != nil{
			return fmt.Errorf("tx err: %v, rbErr: %v", err, rbErr)
		}
	}

	return tx.Commit()
}

// transferTxParams contains the input paraments for a transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID int64 `json:"to_account_id"`
	Amount int64 `json:"amount"`
}

//TransferTxResults is the reult of transfer transaction 
type TransferTxResults struct {
	Transfer Transfer `json:"transfer"`
	FromAccount Account `json:"from_account"`
	ToAccount Account `json:"to_account"`
	FromEntry Entry `json:"from_entry"`
	ToEntry Entry `json:"to_entry"`
}


// for doing money transfer 
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResults, error ) {
	var result TransferTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error 
		
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID: arg.ToAccountID,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID, 
			Amount: -arg.Amount, 
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID, 
			Amount: arg.Amount, 
		})
		if err != nil {
			return err
		}

		// updating balances in both the accounts


		return nil
	})

	return result, err
}