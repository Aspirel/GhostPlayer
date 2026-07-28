package main

import (
    "encoding/json"
    "net/http"
    "time"
    "context"
    "log"
    "go.mongodb.org/mongo-driver/bson"
)

type SearchResult struct {
    Title     string `json:"title"`
    VideoUrl  string `json:"videoUrl"`
    Thumbnail string `json:"thumbnail"`
    Channel   string `json:"channel"`
    Duration  int    `json:"duration"`
}

func YoutubeSearchHandler(w http.ResponseWriter, r *http.Request) {
    collection := DB.Collection("youtube_searches")

    results := []SearchResult{
        {
            Title: "Mock Video 1",
            VideoUrl: "https://www.youtube.com/watch?v=-ncFzlJkJyw",
            Thumbnail: "",
            Channel: "Mock Channel",
            Duration: 120,
        },
        {
            Title: "Mock Video 2",
            VideoUrl: "https://www.youtube.com/watch?v=jBSGD2b2dng",
            Thumbnail: "",
            Channel: "Mock Channel",
            Duration: 240,
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
