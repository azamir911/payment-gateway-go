package data

import "fmt"

type Status string

const (
	StatusNew       Status = "New"
	StatusCompleted Status = "Completed"
	StatusRejected  Status = "Rejected"
)

type card struct {
	Pan    string `json:"pan"`
	Expiry string `json:"expiry"`
}

func (c *card) String() string {
	return fmt.Sprintf("%v %v", c.GetPan(), c.GetExpiry())
}

func (c *card) SetPan(pan string) {
	c.Pan = pan
}

func (c *card) GetPan() string {
	return c.Pan
}

func (c *card) GetExpiry() string {
	return c.Expiry
}

func (c *card) SetExpiry(value string) {
	c.Expiry = value
}

type cardHolder struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (c *cardHolder) GetName() string {
	return c.Name
}

func (c *cardHolder) SetName(value string) {
	c.Name = value
}

func (c *cardHolder) GetEmail() string {
	return c.Email
}

type transaction struct {
	Invoice    int               `json:"invoice"`
	Amount     float64           `json:"amount"`
	Currency   string            `json:"currency"`
	Status     Status            `json:"status"`
	Cardholder *cardHolder       `json:"cardHolder"`
	Card       *card             `json:"card"`
	Errors     map[string]string `json:"errors"`
}

func (t *transaction) String() string {
	return fmt.Sprintf("{%v %v %v %v {%v %v} {%v %v} [%v]", t.GetInvoice(), t.GetAmount(), t.GetCurrency(), t.GetStatus(), t.GetCardHolder().GetName(), t.GetCardHolder().GetEmail(), t.GetCard().GetPan(), t.GetCard().GetExpiry(), t.GetErrors())
}

func (t *transaction) GetInvoice() int {
	return t.Invoice
}

func (t *transaction) GetAmount() float64 {
	return t.Amount
}

func (t *transaction) GetCurrency() string {
	return t.Currency
}

func (t *transaction) GetCardHolder() CardHolder {
	return t.Cardholder
}

func (t *transaction) GetCard() Card {
	return t.Card
}

func (t *transaction) SetStatus(value Status) {
	t.Status = value
}

func (t *transaction) GetStatus() Status {
	return t.Status
}

func (t *transaction) SetErrors(value map[string]string) {
	t.Errors = value
}

func (t *transaction) GetErrors() map[string]string {
	return t.Errors
}
