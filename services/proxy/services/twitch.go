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


type TwitchService struct {
	SearchCollection  *mongo.Collection
	HistoryCollection *mongo.Collection
}

/**
* NewTwitchService creates a new instance of TwitchService
*/
func NewTwitchService(db *mongo.Database) *TwitchService {
	return &TwitchService{
		SearchCollection:  db.Collection("twitch_searches"),
		HistoryCollection: db.Collection("twitch_history"),
	}
}

/**
* SearchHandler handles Twitch search requests
*/
func (ts *TwitchService) SearchHandler(w http.ResponseWriter, r *http.Request) {
    results := []Types.SearchResult{
        {
            Title: "Mock Video 1",
            VideoUrl: "https://www.twitch.tv/videos/2831165142",
            Thumbnail: "",
            Channel: "Mock Channel",
            Duration: 120,
        },
    }

    _, err := ts.SearchCollection.InsertOne(context.TODO(), bson.M{
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
func (ts *TwitchService) GetSearchHistory(w http.ResponseWriter, r *http.Request){
    searchHistory, err := ts.HistoryCollection.Find(context.TODO(), bson.M{})
    if err != nil {
        log.Println("Mongo find failed:", err)
        return
    }

    json.NewEncoder(w).Encode(searchHistory)
}
