package app

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/ent"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
)

type App struct {
	cfg    configs.Config
	logger *logger.Logger
	db     *ent.Client
}

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	db *ent.Client,
) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
		db:     db,
	}
}
