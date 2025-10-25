package main

import (
    "net/http"
    "encoding/json"
    "log"
)

func handlerValidate(w http.ResponseWriter, r *http.Request) {
    type chirpBody struct {
	Body string `json:"body"`
    }
    type chirpError struct {
	Error string `json:"error"`
    }
    type chirpValid struct {
	Valid bool `json:"valid"`
    }

    decoder := json.NewDecoder(r.Body)
    chirp := chirpBody{}
    err := decoder.Decode(&chirp)
    if err != nil {
	log.Printf("Error decoding chirp body: %s", err)
	w.WriteHeader(500)
	return
    }
    if len(chirp.Body) > 140 {
	respBody := chirpError{
	    Error: "Chirp is too long",
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
	    log.Printf("Error marshalling JSON: %s", err)
	    w.WriteHeader(500)
	    return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(400)
	w.Write(dat)
	return
    }

    respBody := chirpValid{
	Valid: true,	
    }

    dat, err := json.Marshal(respBody)
    if err != nil {
	log.Printf("Error marshalling JSON: %s", err)
	w.WriteHeader(500)
	return
    }
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    w.Write(dat)
    return

    
    
}
