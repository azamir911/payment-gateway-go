package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"payment/api"
	"time"

	"payment/data"
	"payment/processor"
	"payment/service"
	"payment/validator"
)

var ts service.TransactionService

func main() {
	newTransactionCh := make(chan data.Transaction, 10)
	savedTransactionCh := make(chan data.Transaction, 10)
	validTransactionCh := make(chan data.Transaction, 10)

	ts = service.GetInstance(newTransactionCh, savedTransactionCh)
	ts.Init()
	vs := validator.GetInstance(savedTransactionCh, validTransactionCh)
	vs.Init()
	ps := processor.GetInstance(validTransactionCh)
	ps.Init()

	fmt.Println("Start running")

	execute()

	api.Serve()

	//log.Info()
	//mux := http.NewServeMux()
	//mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
	//	//rw.Write([]byte("payment gateway started, golang!\n"))
	//	fmt.Fprint(rw, "payment gateway started1, golang!\n")
	//	log.Info().Msg("payment gateway started2, golang!\n")
	//})
	//err := http.ListenAndServe(":3030", mux)
	//if err != nil {
	//	log.Fatal().Msgf("%v", err)
	//}
	//
	//fmt.Println("Running")
}

func execute() {

	//ts := service.GetInstance()

	t := data.NewTransaction(1234567, 10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	//log.Printf("Saving %v", *t)
	log.Logger.Info().Msgf("Saving %v", *t)

	ts.Save(*t)

	time.Sleep(2 * time.Second)

	transaction, _ := ts.Get((*t).GetInvoice())

	//log.Printf("Got %v", transaction)
	log.Logger.Info().Msgf("Got %v", transaction)
}
