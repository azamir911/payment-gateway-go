package validator

import (
	"payment/data"
	"strings"
	"sync"
)

type Valid struct {
	errors map[string]string
}

func NewValid() *Valid {
	var v = Valid{}
	v.errors = make(map[string]string)
	return &v
}

func (v *Valid) IsValid() bool {
	return len(v.errors) == 0
}
func (v *Valid) addError(key string, value string) {
	v.errors[key] = value
}

func (v *Valid) GetError() map[string]string {
	return v.errors
}

type Validator interface {
	Validate(transaction data.Transaction, valid Valid)
}

type allFieldsPresentValidator struct {
}

func (v allFieldsPresentValidator) Validate(transaction data.Transaction, valid Valid) {
	if transaction.GetInvoice() == 0 {
		valid.addError("invoice", "GetInvoice is required")
	}
	if transaction.GetAmount() == 0 {
		valid.addError("amount", "GetAmount is required.")
	}
	if len(strings.TrimSpace(transaction.GetCurrency())) == 0 {
		valid.addError("currency", "GetCurrency is required.")
	}
	if transaction.GetCardHolder() == nil {
		valid.addError("name", "GetName is required.")
		valid.addError("email", "GetEmail is required.")
	} else {
		if len(strings.TrimSpace(transaction.GetCardHolder().GetName())) == 0 {
			valid.addError("name", "GetName is required.")
		}
		if len(strings.TrimSpace(transaction.GetCardHolder().GetEmail())) == 0 {
			valid.addError("email", "GetEmail is required.")
		}
	}
	if transaction.GetCard() == nil {
		valid.addError("pan", "GetPan is required.")
		valid.addError("expiry", "GetExpiry is required.")
	} else {
		if len(strings.TrimSpace(transaction.GetCard().GetPan())) == 0 {
			valid.addError("pan", "GetPan is required.")
		}
		if len(strings.TrimSpace(transaction.GetCard().GetExpiry())) == 0 {
			valid.addError("expiry", "GetExpiry is required.")
		}
	}
}

type positiveAmountValidator struct {
}

func (v positiveAmountValidator) Validate(transaction data.Transaction, valid Valid) {
	if transaction.GetAmount() <= 0 {
		valid.addError("amount", "GetAmount should be a positive double.")
	}
}

var once = sync.Once{}
var instance ValidatorService

type ValidatorService interface {
	Validate(transaction data.Transaction) Valid
}

func GetInstance() ValidatorService {
	once.Do(func() {
		instance = &validatorServiceImpl{}
	})

	return instance
}

type validatorServiceImpl struct {
}

func (v *validatorServiceImpl) Validate(transaction data.Transaction) Valid {
	valid := NewValid()
	a := &allFieldsPresentValidator{}
	a.Validate(transaction, *valid)

	p := &positiveAmountValidator{}
	p.Validate(transaction, *valid)
	return *valid
}
