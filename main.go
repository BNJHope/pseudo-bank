package main

import (
	"database/sql"
	"encoding/json"
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

func getTransactions(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	var (
		transactions []transaction.Transaction
		err          error
	)
	log.Info().Msg("got /transactions request\n")
	w.Header().Set("Content-Type", "application/json")

	transactions, err = tm.GetTransactions()

	if err != nil {
		log.Error().Err(err).Msg("Error getting transactions")
		http.Error(w, "Error fetching transactions", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(transactions)
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

	tm := database.NewPgTransactionManager(db)

	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/transactions", func(w http.ResponseWriter, r *http.Request) {
		getTransactions(w, r, tm)
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Error().Err(err).Msg("Error at root")
	}
}
