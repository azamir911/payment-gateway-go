package data

import "fmt"

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
}

type cardImpl struct {
	pan    string
	expiry string
}

func (c *cardImpl) String() string {
	return fmt.Sprintf("%v %v", c.GetPan(), cardImpl{}.GetExpiry())
}

func (c *cardImpl) SetPan(pan string) {
	c.pan = pan
}

func (c cardImpl) GetPan() string {
	return c.pan
}

func (c cardImpl) GetExpiry() string {
	return c.expiry
}

func (c *cardImpl) SetExpiry(value string) {
	c.expiry = value
}

type cardHolderImpl struct {
	name  string
	email string
}

func (c cardHolderImpl) GetName() string {
	return c.name
}

func (c *cardHolderImpl) SetName(value string) {
	c.name = value
}

func (c cardHolderImpl) GetEmail() string {
	return c.email
}

type transactionImpl struct {
	invoice    int
	amount     float64
	currency   string
	cardholder *cardHolderImpl
	card       *cardImpl
}

func (t transactionImpl) String() string {
	return fmt.Sprintf("{%v %v %v {%v %v} {%v %v}", t.GetInvoice(), t.GetAmount(), t.GetCurrency(), t.GetCardHolder().GetName(), t.GetCardHolder().GetEmail(), t.GetCard().GetPan(), t.GetCard().GetExpiry())
}

func (t transactionImpl) GetInvoice() int {
	return t.invoice
}

func (t transactionImpl) GetAmount() float64 {
	return t.amount
}

func (t transactionImpl) GetCurrency() string {
	return t.currency
}

func (t transactionImpl) GetCardHolder() CardHolder {
	return t.cardholder
}

func (t transactionImpl) GetCard() Card {
	return t.card
}

func NewTransaction(invoice int, amount float64, currency string, name string, email string, pan string, expiry string) *Transaction {

	var transaction Transaction = transactionImpl{
		invoice:    invoice,
		amount:     amount,
		currency:   currency,
		cardholder: &cardHolderImpl{name: name, email: email},
		card:       &cardImpl{pan, expiry},
	}

	return &transaction
}
