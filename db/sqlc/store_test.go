package db

import (
	"context"
	"testing"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testConn)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	// run n concurren transfer trasactions
	n := 5
	amount := int64(10)

	for i := 0; i < n; i++ {
		go func() {
			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: account1.ID,
				ToAccountID:   account2.ID,
				Amount:        amount,
			})
		}()
	}
}
