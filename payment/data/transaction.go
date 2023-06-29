package data

import "fmt"

type Status string

const (
	StatusNew       Status = "New"
	StatusCompleted Status = "Completed"
	StatusRejected  Status = "Rejected"
)

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
