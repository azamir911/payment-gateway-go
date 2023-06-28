package validator

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	transactionRepo "payment/repository"
	"sync"
)

var validators = []validator{&allFieldsPresentValidator{}, &positiveAmountValidator{}}

type Valid struct {
	errors map[string]string
}

func NewValid() *Valid {
	var v = Valid{}
	v.errors = make(map[string]string)
	return &v
}

func (v *Valid) IsValid() bool {
	return v == nil || len(v.errors) == 0
}
func (v *Valid) addError(key string, value string) {
	v.errors[key] = value
}

func (v *Valid) GetErrors() map[string]string {
	return v.errors
}

var once = sync.Once{}
var initOnce = sync.Once{}
var chanIn chan data.Transaction
var chanOut chan<- data.Transaction

var instance ValidatorService

type validatorService struct {
	repo transactionRepo.TransactionRepository
	in   <-chan data.Transaction
	out  chan<- data.Transaction
	done chan struct{}
}

func Init(in chan data.Transaction, out chan<- data.Transaction) {
	initOnce.Do(func() {
		chanIn = in
		chanOut = out
	})
}

func GetInstance() ValidatorService {
	once.Do(func() {
		repository := transactionRepo.GetInstance()
		done := make(chan struct{})
		v := &validatorService{repository, chanIn, chanOut, done}
		instance = v

		go v.init()
	})

	return instance
}

func (v *validatorService) init() {
	for {
		select {
		case transaction := <-v.in:
			log.Logger.Info().Msgf("Got transaction to validate %v", transaction)
			valid := v.Validate(transaction)
			if !valid.IsValid() {
				transaction.SetStatus(data.Status_Rejected)
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
