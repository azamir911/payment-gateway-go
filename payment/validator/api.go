package validator

import "payment/data"

type ValidatorService interface {
	Validate(transaction data.Transaction) Valid
	Close()
}

type validator interface {
	Validate(transaction data.Transaction, valid Valid)
}
