package main

import (
	"ghostplayer/proxy/Services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    Services.InitMongo()

    twitchService := Services.NewTwitchService(Services.DB)
    youtubeService := Services.NewYoutubeService(Services.DB)

    r := mux.NewRouter()
    r.HandleFunc("/youtube/search", youtubeService.SearchHandler).Methods("GET")
    r.HandleFunc("/youtube/history", youtubeService.GetSearchHistory).Methods("GET")
    r.HandleFunc("/twitch/search", twitchService.SearchHandler).Methods("GET")
    r.HandleFunc("/twitch/history", twitchService.GetSearchHistory).Methods("GET")

    log.Println("Proxy running on :8080")
    http.ListenAndServe(":8080", r)
}
