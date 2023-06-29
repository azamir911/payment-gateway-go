package data

type Card interface {
	GetPan() string
	SetPan(value string)
	GetExpiry() string
	SetExpiry(value string)
}

type CardHolder interface {
	GetName() string
	SetName(value string)
	GetEmail() string
}

type Transaction interface {
	GetInvoice() int
	GetAmount() float64
	GetCurrency() string
	GetCardHolder() CardHolder
	GetCard() Card
	SetStatus(value Status)
	GetStatus() Status
	SetErrors(value map[string]string)
	GetErrors() map[string]string
}

func NewEmptyTransaction() *Transaction {
	var transaction Transaction = &transaction{}
	return &transaction
}

func NewTransaction(invoice int, amount float64, currency string, name string, email string, pan string, expiry string) Transaction {

	var transaction Transaction = &transaction{
		Invoice:    invoice,
		Amount:     amount,
		Currency:   currency,
		Cardholder: &cardHolder{Name: name, Email: email},
		Card:       &card{pan, expiry},
		Status:     StatusNew,
		Errors:     nil,
	}

	return transaction
}

func CloneTransaction(t Transaction) *Transaction {

	var transaction Transaction = &transaction{
		Invoice:    t.GetInvoice(),
		Amount:     t.GetAmount(),
		Currency:   t.GetCurrency(),
		Cardholder: &cardHolder{Name: t.GetCardHolder().GetName(), Email: t.GetCardHolder().GetEmail()},
		Card:       &card{t.GetCard().GetPan(), t.GetCard().GetExpiry()},
		Status:     t.GetStatus(),
		Errors:     t.GetErrors(),
	}

	return &transaction
}
