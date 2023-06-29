package processor

import (
	"encoding/base64"
	"payment/data"
	"strings"
)

type panProcessor struct {
}

func (p *panProcessor) Encode(transaction data.Transaction) {
	encode := base64.StdEncoding.EncodeToString([]byte(transaction.GetCard().GetPan()))
	transaction.GetCard().SetPan(encode)
}

func (p *panProcessor) Decode(transaction data.Transaction) {
	decode, _ := base64.StdEncoding.DecodeString(transaction.GetCard().GetPan())
	//TODO check err
	value := string(decode)
	l := len(value)
	value = strings.Repeat("*", l-4) + value[l-4:]
	transaction.GetCard().SetPan(value)
}
