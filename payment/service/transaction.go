package service

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	"payment/processor"
	transactionRepo "payment/repository"
	"sync"
)

var once = sync.Once{}
var initOnce = sync.Once{}
var chanIn chan data.Transaction
var chanOut chan<- data.Transaction

type TransactionService interface {
	Save(transaction data.Transaction)
	Get(invoice int) (data.Transaction, error)
}

var instance TransactionService

type transactionServiceImpl struct {
	repo transactionRepo.TransactionRepository
	in   chan data.Transaction
	out  chan<- data.Transaction
}

func Init(in chan data.Transaction, out chan<- data.Transaction) {
	initOnce.Do(func() {
		chanIn = in
		chanOut = out
	})
}

func GetInstance() TransactionService {
	once.Do(func() {
		repository := transactionRepo.GetInstance()
		t := &transactionServiceImpl{repository, chanIn, chanOut}
		instance = t

		go t.init()
	})

	return instance
}

func (t *transactionServiceImpl) init() {
	for transaction := range t.in {
		log.Logger.Info().Msgf("Got transaction to save %v", transaction)
		t.repo.Save(transaction)
		t.out <- transaction
	}
}

func (t *transactionServiceImpl) Save(transaction data.Transaction) {
	transaction.SetStatus("New")
	t.in <- transaction
}

func (t *transactionServiceImpl) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.repo.Get(invoice)
	if err != nil {
		return nil, err
	}
	if transaction.GetStatus() != "Declined" {
		processor.GetInstance().ApplyDecode(transaction)
	}
	return transaction, err
}
