package main

import (
	"database/sql"
	"io"
	"net/http"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/bnjhope/pseudo-bank/api"
	"github.com/bnjhope/pseudo-bank/database"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("got / request\n")
	io.WriteString(w, "This is my website!\n")
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
	http.HandleFunc("/transaction", func(w http.ResponseWriter, r *http.Request) {
		api.HandleTransaction(w, r, tm)
	})
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		api.HandleUser(w, r, tm)
	})
	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		api.HandleHealthcheck(w, r)
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Error().Err(err).Msg("Error at root")
	}
}
