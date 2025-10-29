package main

import (
    "errors"
    "log"
    "net/http"
)
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, r *http.Request) {
    // Check env variable: PLATFROM == dev
    if cfg.platform != "dev" {
	respondWithError(w, http.StatusForbidden, "Unauthorized environment action", errors.New("UNAUTHORIZED"))
	return
    }

    cfg.fileserverHits.Store(0)
    if err := cfg.db.DeleteAllUsers(r.Context()); err != nil {
	respondWithError(w, http.StatusInternalServerError, "Unable to delete all users", err)
	return
    }
    
    // Response Body
    log.Println("Reset 'Hits' to 0; Deleted all users in database;")
    respondWithJSON(w, http.StatusOK, struct {
	Body	string
    }{
	"Reset Successful",
    })
}
