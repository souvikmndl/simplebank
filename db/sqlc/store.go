package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

/*
This is called composition in Golang. Queries itself does not support Transactions
By embedding it in Store, we allow store to have all the functionalities of Queries
and we can extend Store to have transactions as well
*/

// NewStore creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

// TransferTxParams contains the input params of the transfer tx
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amout         int64 `json:"amount"`
}

// TransferTxResult is the result of the transfer tx
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to the other
// It creates a transfer record, add account entries, and updates accounts' balance in a single db tx
// func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
//     var result TransferTxResult

//     err := store.execTx(ctx, func(q *Queries) error {
//         result.Transfer, err = q.CreateTra
//         return nil
//     })

//     reutrn result, err
// }
