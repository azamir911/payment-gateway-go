package service

import (
	"errors"
	"fmt"
	"payment/data"
	"payment/processor"
	transactionRepo "payment/repository"
	"payment/validator"
	"sync"
)

var once = sync.Once{}

type TransactionService interface {
	Save(transaction data.Transaction) error
	Get(invoice int) (data.Transaction, error)
}

var instance TransactionService

func GetInstance() TransactionService {
	once.Do(func() {
		instance = &transactionServiceImpl{transactionRepo.GetInstance()}
	})

	return instance
}

type transactionServiceImpl struct {
	//var repo = transactionRepo.GetInstance()
	repo transactionRepo.TransactionRepository
}

func (t *transactionServiceImpl) Save(transaction data.Transaction) error {
	validate := validator.GetInstance().Validate(transaction)
	if !validate.IsValid() {
		sprintf := fmt.Sprintf("%v", validate.GetError())
		return errors.New(sprintf)
	}

	processor.GetWriteInstance().ApplyEncode(transaction)
	//defer audit

	return t.repo.Save(transaction)
}

func (t *transactionServiceImpl) Get(invoice int) (data.Transaction, error) {
	transaction, err := t.repo.Get(invoice)
	if err != nil {
		return nil, err
	}
	processor.GetWriteInstance().ApplyDecode(transaction)
	return transaction, err
}
