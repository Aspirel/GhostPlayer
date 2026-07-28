package main

import (
    "encoding/json"
    "net/http"
)

type SearchResult struct {
    Title     string `json:"title"`
    VideoUrl  string `json:"videoUrl"`
    Thumbnail string `json:"thumbnail"`
    Channel   string `json:"channel"`
    Duration  int    `json:"duration"`
}

func YoutubeSearchHandler(w http.ResponseWriter, r *http.Request) {
    //q := r.URL.Query().Get("q")

    // TEMP MOCK — until Node metadata service is ready
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

    json.NewEncoder(w).Encode(results)
}
