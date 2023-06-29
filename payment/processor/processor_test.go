package processor

import (
	"github.com/stretchr/testify/assert"
	"payment/data"
	"testing"
)

func TestPanProcessor(t *testing.T) {
	tr := data.NewTransaction(1234567, 10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	p := &panProcessor{}
	p.Encode(tr)

	assert.Equalf(t, "First Last", (tr).GetCardHolder().GetName(), "name")
	assert.Equalf(t, "NDE4ODg0NjEyMjQ3NjQxMQ==", (tr).GetCard().GetPan(), "Pan should be encrypted")

	p.Decode(tr)
	assert.Equalf(t, "************6411", (tr).GetCard().GetPan(), "Pan should be mask all characters, except last 4 digits")
}

func TestCardHolderProcessor(t *testing.T) {
	tr := data.NewTransaction(1234567, -10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	p := &cardHolderProcessor{}
	p.Encode(tr)

	assert.Equalf(t, "Rmlyc3QgTGFzdA==", (tr).GetCardHolder().GetName(), "Name should be encrypted")

	p.Decode(tr)
	assert.Equalf(t, "**********", (tr).GetCardHolder().GetName(), "Name should be mask all characters")
}

func TestExpiryDateProcessor(t *testing.T) {
	tr := data.NewTransaction(1234567, -10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	p := &expiryDateProcessor{}
	p.Encode(tr)

	assert.Equalf(t, "MDYyNA==", (tr).GetCard().GetExpiry(), "Expiry date should be encrypted")

	p.Decode(tr)
	assert.Equalf(t, "****", (tr).GetCard().GetExpiry(), "Expiry date should be mask all characters")
}
