package processor

import (
	"payment/data"
	transactionRepo "payment/repository"
	"sync"
)

var initOnce = sync.Once{}
var instanceOnce = sync.Once{}

type processor interface {
	Encode(transaction data.Transaction)
	Decode(transaction data.Transaction)
}

type ProcessorRunnner interface {
	ApplyEncode(transaction data.Transaction)
	ApplyDecode(transaction data.Transaction)
	Close()
}

func Init(in chan data.Transaction) {
	initOnce.Do(func() {
		chanIn = in
	})
}

func GetInstance() ProcessorRunnner {
	instanceOnce.Do(func() {
		repository := transactionRepo.GetInstance()
		done := make(chan struct{})
		p := &processorRunner{repository, chanIn, done}
		instance = p

		go p.init()
	})

	return instance
}
