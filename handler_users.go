package main

import (
    "encoding/json"
    "net/http"
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID		uuid.UUID   `json:"id"`
    CreatedAt	time.Time   `json:"created_at"`
    UpdatedAt	time.Time   `json:"updated_at"`
    Email	string	    `json:"email"`
}

func (cfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    type requestBody struct {
	Email string `json:"email"`
    }
    
    decoder := json.NewDecoder(r.Body)
    params := requestBody{}
    if err := decoder.Decode(&params); err != nil {
	respondWithError(w, http.StatusInternalServerError, "Unable to decode request body", err)
	return
    }
    
    // Create new user in database
    user, err := cfg.db.CreateUser(r.Context(), params.Email)
    if err != nil {
	respondWithError(w, http.StatusInternalServerError, "Unable to create new user", err)
	return
    }

    // Response Body
    respondWithJSON(w, http.StatusCreated, User{
	ID: user.ID,
	CreatedAt: user.CreatedAt,
        UpdatedAt: user.UpdatedAt,
	Email: user.Email,
    })
}
