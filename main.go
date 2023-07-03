package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bnjhope/pseudo-bank/database"
	"github.com/bnjhope/pseudo-bank/transaction"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	var (
		transactions []transaction.Transaction
		err          error
	)
	log.Info().Msg("got /transactions request\n")
	w.Header().Set("Content-Type", "application/json")

	transactions, err = database.GetTransactions()

	if err != nil {
		log.Error().Err(err).Msg("Error getting transactions")
		http.Error(w, "Error fetching transactions", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(transactions)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/transactions", getTransactions)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Error().Err(err).Msg("Error at root")
	}
}
