package main

import (
	"context"
	"log"
	"myapp/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func connectMongoDB(uri string) (*mongo.Client, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	client, err := mongo.Connect(opts)

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	return client, nil
}

func main() {
	cfg := config.LoadConfig()

	client, err := connectMongoDB(cfg.MongoURI)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			log.Fatal("Error disconnecting MongoDB:", err)
		}
	}()

	// server := &http.Server{
	// 	Addr:    ":" + cfg.AppPort,
	// 	Handler: nil,
	// }

	// fmt.Println("Starting server on port", cfg.AppPort)
	// if err := server.ListenAndServe(); err != nil {
	// 	log.Fatal("failed to start server: ", err)
	// }

}
