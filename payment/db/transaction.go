package db

import (
	"fmt"
	"payment/data"
	"sync"
)

type TransactionDB struct {
	transactions sync.Map
}

func NewTransactionsDB() *TransactionDB {
	return &TransactionDB{}
}

func (t *TransactionDB) Find(invoice int) (data.Transaction, error) {
	transaction, ok := t.transactions.Load(invoice)
	//transaction, ok := t.transactions[invoice]
	if !ok {
		return nil, fmt.Errorf("invice %v does not exists", invoice)
	}

	return transaction.(data.Transaction), nil
}

func (t *TransactionDB) Upsert(transaction data.Transaction) {
	t.transactions.Store(transaction.GetInvoice(), transaction)
}