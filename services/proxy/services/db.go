package Services

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var DB *mongo.Database

func InitMongo() {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI not set")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    opts := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, opts)

    if err != nil {
        log.Fatal(err)
    }

    // Ping the database to verify the connection
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatal(err)
    }

    Client = client
    DB = client.Database("ghostplayer")

    log.Println("Connected to MongoDB Atlas")
}
