package app

import (
	"github.com/DanielTitkov/tinig-demo-server/internal/configs"
	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/logger"
)

type (
	App struct {
		cfg    configs.Config
		logger *logger.Logger
		repo   Repository
	}
	Repository interface {
		GetUserByUsername(username string) (*domain.User, error)
		CreateUser(user *domain.User) (*domain.User, error)
	}
)

func NewApp(
	cfg configs.Config,
	logger *logger.Logger,
	repo Repository,
) *App {
	return &App{
		cfg:    cfg,
		logger: logger,
		repo:   repo,
	}
}
