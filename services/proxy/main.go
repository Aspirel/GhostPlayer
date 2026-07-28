package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    InitMongo()

    r := mux.NewRouter()
    r.HandleFunc("/youtube/search", YoutubeSearchHandler).Methods("GET")

    log.Println("Proxy running on :8080")
    http.ListenAndServe(":8080", r)
}
