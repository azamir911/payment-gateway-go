package processor

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	transactionRepo "payment/repository"
)

var processors = []processor{&panProcessor{}, &cardHolderProcessor{}, &expiryDateProcessor{}}

var chanIn chan data.Transaction
var instance ProcessorRunnner

type processorRunner struct {
	repo transactionRepo.TransactionRepository
	in   <-chan data.Transaction
	done chan struct{}
}

func (p *processorRunner) init() {
	for {
		select {
		case transaction := <-p.in:
			log.Logger.Info().Msgf("Got transaction to process %v", transaction)
			p.ApplyEncode(transaction)
			transaction.SetStatus(data.StatusCompleted)
			p.repo.Save(transaction)
			log.Logger.Info().Msgf("Got transaction to process2 %v", transaction)
		case <-p.done:
			log.Info().Msg("Transaction processor closed")
			return
		}
	}
}

func (p *processorRunner) ApplyEncode(transaction data.Transaction) {
	for _, p := range processors {
		p.Encode(transaction)
	}
}

func (p *processorRunner) ApplyDecode(transaction data.Transaction) {
	for _, p := range processors {
		p.Decode(transaction)
	}
}

func (p *processorRunner) Close() {
	close(p.done)
}
