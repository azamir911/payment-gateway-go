package service

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	"payment/processor"
	transactionRepo "payment/repository"
	"sync"
)

var once = sync.Once{}

type TransactionService interface {
	Init()
	Save(transaction data.Transaction)
	Get(invoice int) (data.Transaction, error)
}

var instance TransactionService

type transactionServiceImpl struct {
	repo transactionRepo.TransactionRepository
	in   chan data.Transaction
	out  chan<- data.Transaction
}

func GetInstance(in chan data.Transaction, out chan<- data.Transaction) TransactionService {
	once.Do(func() {
		repository := transactionRepo.GetInstance()
		instance = &transactionServiceImpl{repository, in, out}

	})

	return instance
}

func (t *transactionServiceImpl) Init() {
	go func() {
		for transaction := range t.in {
			log.Logger.Info().Msgf("Got transaction to save %v", transaction)
			t.repo.Save(transaction)
			t.out <- transaction
		}
	}()
}

func (t *transactionServiceImpl) Save(transaction data.Transaction) {
	transaction.SetStatus("New")
	t.in <- transaction

	//validate := validator.GetInstance().Validate(transaction)
	//if !validate.IsValid() {
	//	sprintf := fmt.Sprintf("%v", validate.GetError())
	//	return errors.New(sprintf)
	//}
	//
	//processor.GetInstance().ApplyEncode(transaction)
	////defer audit
	//
	//return t.repo.Save(transaction)
}

func (t *transactionServiceImpl) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.repo.Get(invoice)
	if err != nil {
		return nil, err
	}
	if transaction.GetStatus() != "Declined" {
		processor.GetInstance(nil).ApplyDecode(transaction)
	}
	return transaction, err
}
