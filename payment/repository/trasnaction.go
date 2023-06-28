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

type transactionRepository struct {
	transactionsDB *db.TransactionDB
}

func GetInstance() TransactionRepository {
	once.Do(func() {
		transactionsDB := db.NewTransactionsDB()
		instance = &transactionRepository{transactionsDB}
	})

	return instance
}

func (t *transactionRepository) Save(transaction data.Transaction) {
	t.transactionsDB.Upsert(transaction)
}

func (t *transactionRepository) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.transactionsDB.Find(invoice)
	if err != nil {
		return nil, err
	}
	clonedTransaction := data.CloneTransaction(transaction)
	return *clonedTransaction, err
}

func (t *transactionRepository) GetAll() []data.Transaction {
	return t.transactionsDB.FindAll()
}
