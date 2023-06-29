package service

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	"payment/processor"
	transactionRepo "payment/repository"
)

var chanIn chan data.Transaction
var chanOut chan<- data.Transaction

var instance TransactionService

type transactionService struct {
	repo transactionRepo.TransactionRepository
	in   chan data.Transaction
	out  chan<- data.Transaction
	done chan struct{}
}

func (t *transactionService) init() {
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

func (t *transactionService) Save(transaction data.Transaction) {
	transaction.SetStatus(data.StatusNew)
	t.in <- transaction
}

func (t *transactionService) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.repo.Get(invoice)
	if err != nil {
		return nil, err
	}
	if transaction.GetStatus() != data.StatusRejected {
		processor.GetInstance().ApplyDecode(transaction)
	}
	return transaction, err
}

func (t *transactionService) GetAll() []data.Transaction {
	return t.repo.GetAll()
}

func (t *transactionService) Close() {
	close(t.done)
}
