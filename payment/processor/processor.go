package processor

import (
	b64 "encoding/base64"
	"payment/data"
	"sync"
)

var processors = []processor{&panProcessor{}, &cardHolderProcessor{}}

type processor interface {
	Encode(transaction data.Transaction)
	Decode(transaction data.Transaction)
}

type panProcessor struct {
}

func (p *panProcessor) Encode(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetPan()))
	transaction.GetCard().SetPan(encode)
}

func (p *panProcessor) Decode(transaction data.Transaction) {
	decode, _ := b64.StdEncoding.DecodeString(transaction.GetCard().GetPan())
	//TODO check err
	transaction.GetCard().SetPan(string(decode))
}

type cardHolderProcessor struct {
}

func (p *cardHolderProcessor) Encode(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCardHolder().GetName()))
	transaction.GetCardHolder().SetName(encode)
}

func (p *cardHolderProcessor) Decode(transaction data.Transaction) {
	decode, _ := b64.StdEncoding.DecodeString(transaction.GetCardHolder().GetName())
	//TODO check err
	transaction.GetCardHolder().SetName(string(decode))
}

type ProcessorRunnner interface {
	ApplyEncode(transaction data.Transaction)
	ApplyDecode(transaction data.Transaction)
}

var once = sync.Once{}
var instance ProcessorRunnner

type processorRunnerImpl struct {
}

func GetWriteInstance() ProcessorRunnner {
	once.Do(func() {
		instance = &processorRunnerImpl{}
	})

	return instance
}
func (w *processorRunnerImpl) ApplyEncode(transaction data.Transaction) {
	for _, p := range processors {
		p.Encode(transaction)
	}
}

func (w *processorRunnerImpl) ApplyDecode(transaction data.Transaction) {
	for _, p := range processors {
		p.Decode(transaction)
	}
}
