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
	GetAll() []data.Transaction
	Close()
}

var instance TransactionService

type transactionServiceImpl struct {
	repo transactionRepo.TransactionRepository
	in   chan data.Transaction
	out  chan<- data.Transaction
	done chan struct{}
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
		done := make(chan struct{})
		t := &transactionServiceImpl{repository, chanIn, chanOut, done}
		instance = t

		go t.init()
	})

	return instance
}

func (t *transactionServiceImpl) init() {
	for {
		select {
		case transaction := <-t.in:
			log.Logger.Info().Msgf("Got transaction to save %v", transaction)
			t.repo.Save(transaction)
			t.out <- transaction
		case <-t.done:
			log.Info().Msg("Transaction service closed")
			return
		}
	}
}

func (t *transactionServiceImpl) Save(transaction data.Transaction) {
	transaction.SetStatus(data.Status_New)
	t.in <- transaction
}

func (t *transactionServiceImpl) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.repo.Get(invoice)
	if err != nil {
		return nil, err
	}
	if transaction.GetStatus() != data.Status_Rejected {
		processor.GetInstance().ApplyDecode(transaction)
	}
	return transaction, err
}

func (t *transactionServiceImpl) GetAll() []data.Transaction {
	return t.repo.GetAll()
}

func (t *transactionServiceImpl) Close() {
	close(t.done)
}
