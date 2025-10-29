package main

import (
    "database/sql"
    "log"
    "net/http"
    "os"
    "sync/atomic"


    "github.com/joho/godotenv"
    "github.com/kevin-baik/chirpy/internal/database"
    _ "github.com/lib/pq"
)

type apiConfig struct {
    platform	    string
    fileserverHits  atomic.Int32
    db		    *database.Queries
}

func main() {
    const filepathRoot = "."
    const port = "8080"
    
    // Load .env variables
    err := godotenv.Load()
    if err != nil {
	log.Fatal("Error loading .env file")
    }
    
    dbURL := os.Getenv("DB_URL")
    db, err := sql.Open("postgres", dbURL)

    apiCfg := apiConfig{
	platform: os.Getenv("PLATFORM"),
	fileserverHits: atomic.Int32{},
	db: database.New(db),
    }

    mux := http.NewServeMux()

    fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))
    mux.Handle("/app/", fsHandler)

    // API ENDPOINT
    mux.HandleFunc("GET /api/healthz", handlerReadiness)
    mux.HandleFunc("POST /api/validate_chirp", handlerValidate)
    mux.HandleFunc("POST /api/users", apiCfg.handlerCreateUser)
    
    // ADMIN ENDPOINT
    mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
    mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

    srv := &http.Server{
	Addr: ":" + port,
	Handler: mux,
    }

    log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
    log.Fatal(srv.ListenAndServe())

}
