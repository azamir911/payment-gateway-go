package data

import "fmt"

type Status string

const (
	Status_New       Status = "New"
	Status_Completed Status = "Completed"
	Status_Rejected  Status = "Rejected"
)

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

type cardImpl struct {
	Pan    string `json:"pan"`
	Expiry string `json:"expiry"`
}

func (c *cardImpl) String() string {
	return fmt.Sprintf("%v %v", c.GetPan(), cardImpl{}.GetExpiry())
}

func (c *cardImpl) SetPan(pan string) {
	c.Pan = pan
}

func (c cardImpl) GetPan() string {
	return c.Pan
}

func (c cardImpl) GetExpiry() string {
	return c.Expiry
}

func (c *cardImpl) SetExpiry(value string) {
	c.Expiry = value
}

type cardHolderImpl struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c cardHolderImpl) GetName() string {
	return c.Name
}

func (c *cardHolderImpl) SetName(value string) {
	c.Name = value
}

func (c cardHolderImpl) GetEmail() string {
	return c.Email
}

type transaction struct {
	Invoice    int               `json:"invoice"`
	Amount     float64           `json:"amount"`
	Currency   string            `json:"currency"`
	Status     Status            `json:"status"`
	Cardholder *cardHolderImpl   `json:"cardHolder"`
	Card       *cardImpl         `json:"card"`
	Errors     map[string]string `json:"errors"`
}

func (t transaction) String() string {
	return fmt.Sprintf("{%v %v %v %v {%v %v} {%v %v} [%v]", t.GetInvoice(), t.GetAmount(), t.GetCurrency(), t.GetStatus(), t.GetCardHolder().GetName(), t.GetCardHolder().GetEmail(), t.GetCard().GetPan(), t.GetCard().GetExpiry(), t.GetErrors())
}

func (t transaction) GetInvoice() int {
	return t.Invoice
}

func (t transaction) GetAmount() float64 {
	return t.Amount
}

func (t transaction) GetCurrency() string {
	return t.Currency
}

func (t transaction) GetCardHolder() CardHolder {
	return t.Cardholder
}

func (t transaction) GetCard() Card {
	return t.Card
}

func (t *transaction) SetStatus(value Status) {
	t.Status = value
}

func (t transaction) GetStatus() Status {
	return t.Status
}

func (t *transaction) SetErrors(value map[string]string) {
	t.Errors = value
}

func (t transaction) GetErrors() map[string]string {
	return t.Errors
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
		Cardholder: &cardHolderImpl{Name: name, Email: email},
		Card:       &cardImpl{pan, expiry},
		Status:     Status_New,
		Errors:     nil,
	}

	return transaction
}

func CloneTransaction(t Transaction) *Transaction {

	var transaction Transaction = &transaction{
		Invoice:    t.GetInvoice(),
		Amount:     t.GetAmount(),
		Currency:   t.GetCurrency(),
		Cardholder: &cardHolderImpl{Name: t.GetCardHolder().GetName(), Email: t.GetCardHolder().GetEmail()},
		Card:       &cardImpl{t.GetCard().GetPan(), t.GetCard().GetExpiry()},
		Status:     t.GetStatus(),
		Errors:     t.GetErrors(),
	}

	return &transaction
}
