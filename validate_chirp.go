package main

import (
    "net/http"
    "encoding/json"
    "strings"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()
    type requestBody struct {
	Body string `json:"body"`
    }
    type responseBody struct {
	CleanedBody string `json:"cleaned_body"`
    }

    decoder := json.NewDecoder(r.Body)
    params := requestBody{}
    err := decoder.Decode(&params)
    if err != nil {
	respondWithError(w, http.StatusInternalServerError, "Unable to decode request body", err)
	return
    }
    const maxChirpLength = 140
    if len(params.Body) > maxChirpLength {
	respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
	return
    }
    
    respondWithJSON(w, http.StatusOK, responseBody{
	CleanedBody: replaceBadWords(params.Body),
    })

}

func replaceBadWords(str string) string {
    badWords := map[string]bool{
	"kerfuffle": true,
	"sharbert": true,
	"fornax": true,
    }
    strSplit := strings.Split(str, " ")
    for i, word := range strSplit {
	if badWords[strings.ToLower(word)] {
	    strSplit[i] = "****"
	}
    }
    return strings.Join(strSplit, " ")
}
