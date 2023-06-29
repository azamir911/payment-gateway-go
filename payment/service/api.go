package service

import (
	"payment/data"
	transactionRepo "payment/repository"
	"sync"
)

var initOnce = sync.Once{}
var instanceOnce = sync.Once{}

type TransactionService interface {
	Save(transaction data.Transaction)
	Get(invoice int) (data.Transaction, error)
	GetAll() []data.Transaction
	Close()
}

func Init(in chan data.Transaction, out chan<- data.Transaction) {
	initOnce.Do(func() {
		chanIn = in
		chanOut = out
	})
}

func GetInstance() TransactionService {
	instanceOnce.Do(func() {
		repository := transactionRepo.GetInstance()
		done := make(chan struct{})
		t := &transactionService{repository, chanIn, chanOut, done}
		instance = t

		go t.init()
	})

	return instance
}
