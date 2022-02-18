package processor

import (
	b64 "encoding/base64"
	"payment/data"
	"sync"
)

type Processor interface {
	Encode(transaction data.Transaction)
	Decode(transaction data.Transaction)
}

type PanProcessor struct {
}

func (p *PanProcessor) Encode(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetPan()))
	transaction.GetCard().SetPan(encode)
}

func (p *PanProcessor) Decode(transaction data.Transaction) {
	decode, _ := b64.StdEncoding.DecodeString(transaction.GetCard().GetPan())
	//TODO check err
	transaction.GetCard().SetPan(string(decode))
}

type CardHolderProcessor struct {
}

func (p *CardHolderProcessor) Encode(transaction data.Transaction) {
	encode := b64.StdEncoding.EncodeToString([]byte(transaction.GetCardHolder().GetName()))
	transaction.GetCardHolder().SetName(encode)
}

func (p *CardHolderProcessor) Decode(transaction data.Transaction) {
	decode, _ := b64.StdEncoding.DecodeString(transaction.GetCardHolder().GetName())
	//TODO check err
	transaction.GetCardHolder().SetName(string(decode))
}

type ProcessorRunnner interface {
	ApplyEncode(transaction data.Transaction)
	ApplyDecode(transaction data.Transaction)
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
func (w *writeProcessorRunnerImpl) ApplyEncode(transaction data.Transaction) {
	p := &PanProcessor{}
	c := &CardHolderProcessor{}
	p.Encode(transaction)
	c.Encode(transaction)
}

func (w *writeProcessorRunnerImpl) ApplyDecode(transaction data.Transaction) {
	p := &PanProcessor{}
	c := &CardHolderProcessor{}
	p.Decode(transaction)
	c.Decode(transaction)
}

//var readOnce = sync.Once{}
//
//var readInstance ProcessorRunnner
//
//type readProcessorRunnerImpl struct {
//}
//
//func GetReadInstance() ProcessorRunnner {
//	readOnce.Do(func() {
//		readInstance = &readProcessorRunnerImpl{}
//	})
//
//	return readInstance
//}
//
//func (r *readProcessorRunnerImpl) ApplyEncode(transaction data.Transaction) {
//	//TODO implement me
//	panic("implement me")
//}
