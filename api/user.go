package api

import (
	"fmt"
	"net/http"

	"github.com/bnjhope/pseudo-bank/database"
	"github.com/rs/zerolog/log"
)

func HandleUser(w http.ResponseWriter, r *http.Request, tm database.TransactionManager) {
	log.Error().Err(fmt.Errorf("unrecognised request method %v", r.Method)).Msg("Error handling /user request")
	http.Error(w, "Error fetching users", http.StatusInternalServerError)
}
