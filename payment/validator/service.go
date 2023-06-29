package validator

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	transactionRepo "payment/repository"
)

var validators = []validator{&allFieldsPresentValidator{}, &positiveAmountValidator{}}

var chanIn chan data.Transaction
var chanOut chan<- data.Transaction

var instance ValidatorService

type validatorService struct {
	repo transactionRepo.TransactionRepository
	in   <-chan data.Transaction
	out  chan<- data.Transaction
	done chan struct{}
}

func (v *validatorService) init() {
	for {
		select {
		case transaction := <-v.in:
			log.Logger.Info().Msgf("Got transaction to validate %v", transaction)
			valid := v.Validate(transaction)
			if !valid.IsValid() {
				transaction.SetStatus(data.StatusRejected)
				transaction.SetErrors(valid.GetErrors())
				v.repo.Save(transaction)
			} else {
				v.out <- transaction
			}
		case <-v.done:
			log.Info().Msg("Transaction validator closed")
			return
		}
	}
}

func (v *validatorService) Validate(transaction data.Transaction) Valid {
	valid := NewValid()
	for _, v := range validators {
		v.Validate(transaction, *valid)
	}

	return *valid
}

func (v *validatorService) Close() {
	close(v.done)
}
