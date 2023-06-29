package repository

import (
	"payment/data"
	"payment/db"
	"sync"
)

var instanceOnce = sync.Once{}

type TransactionRepository interface {
	Save(transaction data.Transaction)
	Get(invoice int) (data.Transaction, error)
	GetAll() []data.Transaction
}

func GetInstance() TransactionRepository {
	instanceOnce.Do(func() {
		transactionsDB := db.NewTransactionsDB()
		instance = &transactionRepository{transactionsDB}
	})

	return instance
}
