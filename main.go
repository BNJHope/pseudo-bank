package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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

func handleTransaction(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
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

	transactions, err := tm.GetTransactions()

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

	if err != nil {
		log.Error().Err(err).Msg("Error saving transaction")
		http.Error(w, "Error saving transaction", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(struct {
		Result int64 `json:"transactionId"`
	}{transactionId})
}

func setupDB() (*sql.DB, error) {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	db, dbSetupErr := setupDB()

	if dbSetupErr != nil {
		log.Error().Err(dbSetupErr).Msg("Error setting up database")
	}

	defer db.Close()

	tm := database.NewPgTransactionManager(db)

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		handleTransaction(w, r, tm)
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Error().Err(err).Msg("Error at root")
	}
}
