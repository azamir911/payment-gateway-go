package repository

import (
	"payment/data"
	"payment/db"
	"sync"
)

var once = sync.Once{}

type TransactionRepository interface {
	Save(transaction data.Transaction)
	Get(invoice int) (data.Transaction, error)
}

var instance TransactionRepository

type transactionRepositoryImpl struct {
	transactionsDB *db.TransactionDB
}

func GetInstance() TransactionRepository {
	once.Do(func() {
		transactionsDB := db.NewTransactionsDB()
		instance = &transactionRepositoryImpl{transactionsDB}
	})

	return instance
}

func (t *transactionRepositoryImpl) Save(transaction data.Transaction) {
	t.transactionsDB.Upsert(transaction)
}

func (t *transactionRepositoryImpl) Get(invoice int) (data.Transaction, error) {
	return t.transactionsDB.Find(invoice)
}
