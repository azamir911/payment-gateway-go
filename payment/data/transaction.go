package data

type Card interface {
	Pan() string
	Expiry() string
}

type CardHolder interface {
	Name() string
	Email() string
}

type Transaction interface {
	Invoice() int
	Amount() float64
	Currency() string
	CardHolder() CardHolder
	Card() Card
}

type cardImpl struct {
	pan    string
	expiry string
}

func (c cardImpl) Pan() string {
	return c.pan
}

func (c cardImpl) Expiry() string {
	return c.expiry
}

type cardHolderImpl struct {
	name  string
	email string
}

func (c cardHolderImpl) Name() string {
	return c.name
}

func (c cardHolderImpl) Email() string {
	return c.email
}

type transactionImpl struct {
	invoice    int
	amount     float64
	currency   string
	cardholder cardHolderImpl
	card       cardImpl
}

func (t transactionImpl) Invoice() int {
	return t.invoice
}

func (t transactionImpl) Amount() float64 {
	return t.amount
}

func (t transactionImpl) Currency() string {
	return t.currency
}

func (t transactionImpl) CardHolder() CardHolder {
	return t.cardholder
}

func (t transactionImpl) Card() Card {
	return t.card
}

func NewTransaction(invoice int, amount float64, currency string, name string, email string, pan string, expiry string) *Transaction {

	//CardImpl := &CardImpl{pan: pan, expiry: expiry}
	//cardHolder := &CardHolderImpl{
	//	Name:  Name,
	//	Email: Email,
	//}
	var transaction Transaction = transactionImpl{
		invoice:    invoice,
		amount:     amount,
		currency:   currency,
		cardholder: cardHolderImpl{name: name, email: email},
		card:       cardImpl{pan, expiry},
	}

	return &transaction
}
