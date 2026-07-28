package main

import (
    "context"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func InitMongo() {
    uri := os.Getenv("MONGO_URI")
    if uri == "" {
        log.Fatal("MONGO_URI not set")
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    DB = client.Database("ghostplayer")

    log.Println("Connected to MongoDB Atlas")
}
