package main

import (
	"fmt"
	_ "github.com/qodrorid/godaemon"
	"github.com/rs/zerolog/log"
	"net/http"

	"payment/data"
	"payment/processor"
	"payment/service"
	"payment/validator"
	"time"
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

	execute()

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("Assalamu'alaikum, golang!\n"))
	})
	log.Fatal().Msgf("%v", http.ListenAndServe(":3030", mux))
	//log.Fatalln(http.ListenAndServe(":3030", mux))
}

func execute() {
	fmt.Println("Start running")

	//ts := service.GetInstance()

	t := data.NewTransaction(1234567, -10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	//log.Printf("Saving %v", *t)
	log.Logger.Info().Msgf("Saving %v", *t)

	ts.Save(*t)

	time.Sleep(2 * time.Second)

	transaction, _ := ts.Get((*t).GetInvoice())

	//log.Printf("Got %v", transaction)
	log.Logger.Info().Msgf("Got %v", transaction)
}
