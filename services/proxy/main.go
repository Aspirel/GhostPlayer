package main

import (
	"ghostplayer/proxy/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    services.InitMongo()

    r := mux.NewRouter()
    r.HandleFunc("/youtube/search", services.YoutubeSearchHandler).Methods("GET")
    r.HandleFunc("/twitch/search", services.TwitchSearchHandler).Methods("GET")

    log.Println("Proxy running on :8080")
    http.ListenAndServe(":8080", r)
}
