package repository

import (
	"payment/data"
	"payment/db"
)

var instance TransactionRepository

type transactionRepository struct {
	transactionsDB *db.TransactionDB
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
