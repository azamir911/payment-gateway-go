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
	GetAll() []data.Transaction
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
	transaction, err := t.transactionsDB.Find(invoice)
	if err != nil {
		return nil, err
	}
	clonedTransaction := data.CloneTransaction(transaction)
	return *clonedTransaction, err
}

func (t *transactionRepositoryImpl) GetAll() []data.Transaction {
	return t.transactionsDB.FindAll()
}
