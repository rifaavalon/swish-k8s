package main

import (
    "log"
    "net/http"
    "go-app/handlers"
    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/api/environments", handlers.CreateEnvironment).Methods("POST")

    log.Println("Starting server on :8080...")
    if err := http.ListenAndServe(":8080", r); err != nil {
        log.Fatalf("Server failed: %s", err)
    }
}