package main

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"payment/data"
	"payment/service"
)

func main() {
	fmt.Println("Start running")

	ts := service.GetInstance()

	t := data.NewTransaction(1234567, 10, "EUR", "First Last", "email@domain.com", "4188846122476411", "0624")

	log.Logger.Info().Msgf("Saving %v", *t)

	err := ts.Save(*t)

	if err != nil {
		log.Logger.Fatal().Msgf("Error while trying to save transaction %v, err is %v", *t, err)
	}

	log.Logger.Info().Msgf("Saved %v", *t)

	transaction, err := ts.Get((*t).GetInvoice())

	log.Logger.Info().Msgf("Got %v", transaction)
}
