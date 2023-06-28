package validator

import (
	"payment/data"
	"strings"
)

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
