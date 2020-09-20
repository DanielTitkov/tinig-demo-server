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
		GetTaskByCode(code string) (*domain.Task, error)
		GetTaskTypeByCode(code string) (*domain.TaskType, error)
		CreateUser(*domain.User) (*domain.User, error)
		CreateTask(*domain.Task, *domain.User, *domain.TaskType) (*domain.Task, error)
		CreateItem(*domain.Item, *domain.Task) (*domain.Item, error)
		GetTasks(u *domain.User, itemLimit int, deactivated bool) ([]*domain.Task, error)
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
