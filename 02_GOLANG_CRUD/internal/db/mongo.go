package db

import (
	"context"
	"fmt"
	"notes-api/internal/config"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	// clientOpts := options.Client().ApplyURI(cfg.MongoURI)
// 	clientOpts := options.Client().ApplyURI(cfg.MongoURI)

// 	// client, err := mongo.Connect(clientOpts)
// 	client, err := mongo.NewClient(clientOpts)
// 	if err != nil {
// 		return nil, nil, fmt.Errorf("mongo connect failed")
// 	}

// 	if err := client.Connect(ctx); err != nil {
//     return nil, nil, fmt.Errorf("mongo connect failed: %v", err)
// 	}

// 	if err := client.Ping(ctx, nil); err != nil {
// 		return nil, nil, fmt.Errorf("mongo ping failed")
// 	}

// 	database := client.Database(cfg.MongoDB)
// 	return client, database, nil
// }

func Connect(cfg config.Config) (*mongo.Client, *mongo.Database, error) {

	// Now use ctx when you actually talk to the server
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Build client options
	clientOpts := options.Client().ApplyURI(cfg.MongoURI)

	// Create client (no ctx here)
	client, err := mongo.Connect(clientOpts)
	if err != nil {
		return nil, nil, fmt.Errorf("mongo connect failed: %v", err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("mongo ping failed: %v", err)
	}

	database := client.Database(cfg.MongoDB)
	return client, database, nil
}

func Disconnect(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return client.Disconnect(ctx)
}
