package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bnjhope/pseudo-bank/database"
	"github.com/rs/zerolog/log"
)

func HandleUser(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	if r.Method == http.MethodGet {
		getUser(w, r, tm)
	} else {
		log.Error().Err(fmt.Errorf("unrecognised request method %v", r.Method)).Msg("Error handling /user request")
		http.Error(w, "Error fetching users", http.StatusInternalServerError)
	}
}

func getUser(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	log.Info().Msg("got /user request\n")
	w.Header().Set("Content-Type", "application/json")

	userId := r.URL.Query().Get("id")

	user, err := tm.GetUser(userId)

	if err != nil {
		log.Error().Err(err).Msg("Error getting user")
		http.Error(w, "Error fetching user", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(user)
}
