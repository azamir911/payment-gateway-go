package processor

import (
	"encoding/base64"
	"payment/data"
	"strings"
)

type expiryDateProcessor struct {
}

func (e expiryDateProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetExpiry()))
	transaction.GetCard().SetExpiry(encode)
}

func (e expiryDateProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCard().GetExpiry())
	//TODO check err
	l := len(string(decode))
	value := strings.Repeat("*", l)
	transaction.GetCard().SetExpiry(value)
}
