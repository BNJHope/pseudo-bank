package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/bnjhope/pseudo-bank/database"
	"github.com/bnjhope/pseudo-bank/transaction"
	"github.com/rs/zerolog/log"
)

func HandleTransaction(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	if r.Method == http.MethodGet {
		getTransactions(w, r, tm)
	} else if r.Method == http.MethodPost {
		addTransaction(w, r, tm)
	} else {
		log.Error().Err(fmt.Errorf("unrecognised request method %v", r.Method)).Msg("Error handling /transaction request")
		http.Error(w, "Error fetching transactions", http.StatusInternalServerError)
	}
}

func getTransactions(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	log.Info().Msg("got /transaction request\n")
	w.Header().Set("Content-Type", "application/json")

	userId := r.URL.Query().Get("from")

	transactions, err := tm.GetTransactions(userId)

	if err != nil {
		log.Error().Err(err).Msg("Error getting transactions")
		http.Error(w, "Error fetching transactions", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(transactions)
}

func addTransaction(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	log.Info().Msg("got post /transaction request\n")
	w.Header().Set("Content-Type", "application/json")

	var t transaction.Transaction

	json.NewDecoder(r.Body).Decode(&t)

	transactionId, err := tm.SaveTransaction(&t)

	target := &transaction.NotEnoughFundsInAccountError{}
	if errors.As(err, &target) {
		serr, _ := err.(*transaction.NotEnoughFundsInAccountError)
		log.Error().Err(serr).Msg("Not enough funds in account")
		http.Error(w, serr.Error(), http.StatusBadRequest)
	} else if err != nil {
		log.Error().Err(err).Msg("Error saving transaction")
		http.Error(w, "Error saving transaction", http.StatusInternalServerError)
	} else {
		json.NewEncoder(w).Encode(struct {
			Result int64 `json:"transactionId"`
		}{transactionId})
	}
}
