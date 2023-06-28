package validator

import "payment/data"

type positiveAmountValidator struct {
}

func (v positiveAmountValidator) Validate(transaction data.Transaction, valid Valid) {
	if transaction.GetAmount() <= 0 {
		valid.addError("amount", "Amount should be a positive double.")
	}
}
