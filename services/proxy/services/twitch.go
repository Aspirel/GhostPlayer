package services

import (
	"context"
	"encoding/json"
	"ghostplayer/proxy/types"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func TwitchSearchHandler(w http.ResponseWriter, r *http.Request) {
    collection := DB.Collection("twitch_searches")

    results := []types.SearchResult{
        {
            Title: "Mock Video 1",
            VideoUrl: "https://www.twitch.tv/videos/2831165142",
            Thumbnail: "",
            Channel: "Mock Channel",
            Duration: 120,
        },
    }

    _, err := collection.InsertOne(context.TODO(), bson.M{
        "timestamp": time.Now(),
        "results":   results,
    })

    if err != nil {
        log.Println("Mongo insert failed:", err)
    }

    json.NewEncoder(w).Encode(results)
}
