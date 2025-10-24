package main

import (
    //"sync/atomic"
    "io"
    "log"
    "net/http"
//    "fmt"

)

func main() {
    const filepathRoot = "."
    const port = "8080"
    
    /*
    apiCfg := apiConfig{
	fileserverHits: atomic.Int32{},
    }
    */

    mux := http.NewServeMux()
    mux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
    mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "OK")
    })
    //mux.Handle("/metrics", apiCfg.handlerMetrics(func(w http.ResponseWriter, req *http.Request)))

    server := http.Server{
	Addr: ":" + port,
	Handler: mux,
    }

    log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
    log.Fatal(server.ListenAndServe())

}

/*
type apiConfig struct {
    fileserverHits atomic.Int32
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
    cfg.fileserverHits.Add(1)
    return next
}

func (cfg *apiConfig) handlerMetrics(func(w http.ResponseWriter, req *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, req *http.Request){
	    w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	    w.WriteHeader(http.StatusOK)
	    io.WriteString(w, fmt.Sprintf("Hits: %v", cfg.fileserverHits))
	}
}
*/

