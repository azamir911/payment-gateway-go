package processor

import (
	b64 "encoding/base64"
	"payment/data"
	"sync"
)

type Processor interface {
	Process(transaction data.Transaction)
}

type EncodePanProcessor struct {
}

func (p *EncodePanProcessor) Process(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetPan()))
	transaction.GetCard().SetPan(encode)
}

type EncodeCardHolderProcessor struct {
}

func (p *EncodeCardHolderProcessor) Process(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCardHolder().GetName()))
	transaction.GetCardHolder().SetName(encode)
}

type ProcessorRunnner interface {
	Apply(transaction data.Transaction)
}

var writeOnce = sync.Once{}
var writeInstance ProcessorRunnner

type writeProcessorRunnerImpl struct {
}

func GetWriteInstance() ProcessorRunnner {
	writeOnce.Do(func() {
		writeInstance = &writeProcessorRunnerImpl{}
	})

	return writeInstance
}
func (w *writeProcessorRunnerImpl) Apply(transaction data.Transaction) {
	p := &EncodePanProcessor{}
	c := &EncodeCardHolderProcessor{}
	p.Process(transaction)
	c.Process(transaction)
}

var readOnce = sync.Once{}

var readInstance ProcessorRunnner

type readProcessorRunnerImpl struct {
}

func GetReadInstance() ProcessorRunnner {
	readOnce.Do(func() {
		readInstance = &readProcessorRunnerImpl{}
	})

	return readInstance
}

func (r *readProcessorRunnerImpl) Apply(transaction data.Transaction) {
	//TODO implement me
	panic("implement me")
}


