package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "ghostplayer/proxy/services"
)

func main() {
    services.InitMongo()

    r := mux.NewRouter()
    r.HandleFunc("/youtube/search", services.YoutubeSearchHandler).Methods("GET")

    log.Println("Proxy running on :8080")
    http.ListenAndServe(":8080", r)
}
