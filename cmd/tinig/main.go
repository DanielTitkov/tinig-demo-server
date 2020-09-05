package main

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/cmd/tinig/prepare"
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/ent"

	_ "github.com/lib/pq"
)

func main() {
	logger := prepare.NewLogger()
	defer logger.Sync()
	logger.Info("starting service", "")

	cfg, err := configs.ReadConfigs("./configs/dev.yml")
	if err != nil {
		logger.Fatal("failed to load config", err)
	}
	logger.Info("loaded config", "")

	client, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer client.Close()
	logger.Info("connected to database", cfg.DB.Driver+"://"+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), client) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")
}
