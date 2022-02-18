package repository

import (
	"fmt"
	"payment/data"
	"sync"
)

var once = sync.Once{}

type TransactionRepository interface {
	Save(transaction data.Transaction) error
	Get(invoice int) (data.Transaction, error)
}

var instance TransactionRepository

func GetInstance() TransactionRepository {
	once.Do(func() {
		instance = &transactionRepositoryImpl{}
	})

	return instance
}

type transactionRepositoryImpl struct {
	//transactions map[int]data.Transaction
	transactions sync.Map
}

func (t *transactionRepositoryImpl) Save(transaction data.Transaction) error {
	//t.transactions[transaction.GetInvoice()] = transaction
	t.transactions.Store(transaction.GetInvoice(), transaction)

	return nil
}

func (t *transactionRepositoryImpl) Get(invoice int) (data.Transaction, error) {
	transaction, ok := t.transactions.Load(invoice)
	//transaction, ok := t.transactions[invoice]
	if !ok {
		return nil, fmt.Errorf("invice %v does not exists", invoice)
	}

	return transaction.(data.Transaction), nil

}
