package main

import (
    "net/http"
    "log"

)

func main() {
    mux := http.NewServeMux()
    mux.Handle("/", http.FileServer(http.Dir(".")))
    server := http.Server{
	Addr: ":8080",
	Handler: mux,
    }
    
    err := server.ListenAndServe()
    if err != nil {
	log.Fatalf("ListenAndServe Error: %w", err)
    }

}
