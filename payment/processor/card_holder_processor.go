package processor

import (
	"encoding/base64"
	"payment/data"
	"strings"
)

type cardHolderProcessor struct {
}

func (p *cardHolderProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCardHolder().GetName()))
	transaction.GetCardHolder().SetName(encode)
}

func (p *cardHolderProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCardHolder().GetName())
	//TODO check err
	l := len(string(decode))
	value := strings.Repeat("*", l)
	transaction.GetCardHolder().SetName(value)
}
