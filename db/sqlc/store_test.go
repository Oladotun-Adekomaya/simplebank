package db

import "testing"

func TestTransferTx(t *testing.T) {
	store := NewStore(testConn)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
}
