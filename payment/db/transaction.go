package db

import (
	"fmt"
	"payment/data"
	"sync"
)

type TransactionDB struct {
	transactions sync.Map
}

func (t *TransactionDB) Find(invoice int) (data.Transaction, error) {
	transaction, ok := t.transactions.Load(invoice)
	if !ok {
		return nil, fmt.Errorf("invice %v does not exists", invoice)
	}

	return transaction.(data.Transaction), nil
}

func (t *TransactionDB) Upsert(transaction data.Transaction) {
	t.transactions.Store(transaction.GetInvoice(), transaction)
}

func (t *TransactionDB) FindAll() []data.Transaction {
	var all []data.Transaction
	t.transactions.Range(func(_, value interface{}) bool {
		all = append(all, value.(data.Transaction))
		return true
	})
	return all
}
