package Services

import (
	"context"
	"encoding/json"
	"ghostplayer/proxy/Types"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type YoutubeService struct {
	SearchCollection  *mongo.Collection
	HistoryCollection *mongo.Collection
}

/**
* NewYoutubeService creates a new instance of YoutubeService
*/
func NewYoutubeService(db *mongo.Database) *YoutubeService {
	return &YoutubeService{
		SearchCollection:  db.Collection("youtube_searches"),
		HistoryCollection: db.Collection("youtube_history"),
	}
}

/**
* SearchHandler handles YouTube search requests
*/
func (ys *YoutubeService) SearchHandler(w http.ResponseWriter, r *http.Request) {
    results := []Types.SearchResult{
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

    _, err := ys.SearchCollection.InsertOne(context.TODO(), bson.M{
        "timestamp": time.Now(),
        "results":   results,
    })

    if err != nil {
        log.Println("Mongo insert failed:", err)
    }

    json.NewEncoder(w).Encode(results)
}

/**
* GetSearchHistory retrieves the search history from the database
*/
func (ys *YoutubeService) GetSearchHistory(w http.ResponseWriter, r *http.Request){
    searchHistory, err := ys.HistoryCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Println("Mongo find failed:", err)
        return
    }

    json.NewEncoder(w).Encode(searchHistory)
}
