package processor

import (
	"encoding/base64"
	"payment/data"
	"strings"
	"sync"
)

var processors = []processor{&panProcessor{}, &cardHolderProcessor{}, &expiryDateProcessor{}}

type processor interface {
	Encode(transaction data.Transaction)
	Decode(transaction data.Transaction)
}

type panProcessor struct {
}

func (p *panProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetPan()))
	transaction.GetCard().SetPan(encode)
}

func (p *panProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCard().GetPan())
	//TODO check err
	value := string(decode)
	l := len(value)
	value = strings.Repeat("*", l-4) + value[l-4:]
	transaction.GetCard().SetPan(value)
}

type expiryDateProcessor struct {
}

func (e expiryDateProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetExpiry()))
	transaction.GetCard().SetExpiry(encode)
}

func (e expiryDateProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCard().GetExpiry())
	//TODO check err
	l := len(string(decode))
	value := strings.Repeat("*", l)
	transaction.GetCard().SetExpiry(value)
}

type cardHolderProcessor struct {
}

func (p *cardHolderProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCardHolder().GetName()))
	transaction.GetCardHolder().SetName(encode)
}

func (p *cardHolderProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCardHolder().GetName())
	//TODO check err
	l := len(string(decode))
	value := strings.Repeat("*", l)
	transaction.GetCardHolder().SetName(value)
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
