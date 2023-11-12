package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
)

func HandleHealthcheck(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		healthcheckSuccess(w, r)
	} else {
		log.Error().Err(fmt.Errorf("unrecognised request method %v", r.Method)).Msg("Error handling /healthcheck request")
		http.Error(w, "Error handling healthcheck", http.StatusInternalServerError)
	}
}

func healthcheckSuccess(w http.ResponseWriter, r *http.Request) {
	log.Info().Msg("got /healthcheck request\n")
	w.Header().Set("Content-Type", "application/json")
	success_response := struct {
		Success bool `json:"success"`
	}{true}

	json.NewEncoder(w).Encode(success_response)
}
