package app

import (
	"context"
	"fmt"
	"go-auth/internal/config"
	"go-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New() (*App, error) {
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}

	mongoCli, err := db.Connect(cfg)

	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.DB,
	}, nil

}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}

	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("mongo disconnect failed %w", err)
	}

	return nil

}
