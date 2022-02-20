package validator

import (
	"github.com/rs/zerolog/log"
	"payment/data"
	transactionRepo "payment/repository"
	"strings"
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

type validator interface {
	Validate(transaction data.Transaction, valid Valid)
}

type allFieldsPresentValidator struct {
}

func (v allFieldsPresentValidator) Validate(transaction data.Transaction, valid Valid) {
	if transaction.GetInvoice() == 0 {
		valid.addError("invoice", "Invoice is required")
	}
	if transaction.GetAmount() == 0 {
		valid.addError("amount", "Amount is required.")
	}
	if len(strings.TrimSpace(transaction.GetCurrency())) == 0 {
		valid.addError("currency", "Currency is required.")
	}
	if transaction.GetCardHolder() == nil {
		valid.addError("name", "Name is required.")
		valid.addError("email", "Email is required.")
	} else {
		if len(strings.TrimSpace(transaction.GetCardHolder().GetName())) == 0 {
			valid.addError("name", "Name is required.")
		}
		if len(strings.TrimSpace(transaction.GetCardHolder().GetEmail())) == 0 {
			valid.addError("email", "Email is required.")
		}
	}
	if transaction.GetCard() == nil {
		valid.addError("pan", "Pan is required.")
		valid.addError("expiry", "Expiry is required.")
	} else {
		if len(strings.TrimSpace(transaction.GetCard().GetPan())) == 0 {
			valid.addError("pan", "Pan is required.")
		}
		if len(strings.TrimSpace(transaction.GetCard().GetExpiry())) == 0 {
			valid.addError("expiry", "Expiry is required.")
		}
	}
}

type positiveAmountValidator struct {
}

func (v positiveAmountValidator) Validate(transaction data.Transaction, valid Valid) {
	if transaction.GetAmount() <= 0 {
		valid.addError("amount", "Amount should be a positive double.")
	}
}

var once = sync.Once{}
var initOnce = sync.Once{}
var chanIn chan data.Transaction
var chanOut chan<- data.Transaction

var instance ValidatorService

type ValidatorService interface {
	Validate(transaction data.Transaction) Valid
	Close()
}

type validatorServiceImpl struct {
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
		v := &validatorServiceImpl{repository, chanIn, chanOut, done}
		instance = v

		go v.init()
	})

	return instance
}

func (v *validatorServiceImpl) init() {
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

func (v *validatorServiceImpl) Validate(transaction data.Transaction) Valid {
	valid := NewValid()
	for _, v := range validators {
		v.Validate(transaction, *valid)
	}

	return *valid
}

func (v *validatorServiceImpl) Close() {
	close(v.done)
}
