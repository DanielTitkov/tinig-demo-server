package main

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/cmd/tinig/prepare"
	"github.com/DanielTitkov/tinig-demo-server/internal/app"
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/job"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent"

	_ "github.com/lib/pq"
)

func main() {
	logger := logger.NewLogger()
	defer logger.Sync()
	logger.Info("starting service", "")

	cfg, err := configs.ReadConfigs("./configs/dev.yml")
	if err != nil {
		logger.Fatal("failed to load config", err)
	}
	logger.Info("loaded config", "")

	db, err := ent.Open(cfg.DB.Driver, cfg.DB.URI)
	if err != nil {
		logger.Fatal("failed connecting to database", err)
	}
	defer db.Close()
	logger.Info("connected to database", cfg.DB.Driver+", "+cfg.DB.URI)

	err = prepare.Migrate(context.Background(), db) // run db migration
	if err != nil {
		logger.Fatal("failed creating schema resources", err)
	}
	logger.Info("migrations done", "")

	repo := entgo.NewEntgoRepository(db, logger)

	app := app.NewApp(cfg, logger, repo)

	jobs := job.NewService(cfg, logger, app)
	jobs.GatherSystemSummary() // TODO: maybe hide it inside jobs

	server := prepare.NewServer(cfg, logger, app)
	logger.Fatal("failed to start server", server.Start(cfg.Server.GetAddress()))
}
