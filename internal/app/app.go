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
		// users
		CreateUser(*domain.User) (*domain.User, error)
		GetUserByUsername(username string) (*domain.User, error)
		GetUserCount() (int, error)

		// tasks
		CreateTask(*domain.Task, *domain.User, *domain.TaskType) (*domain.Task, error)
		GetTasks(u *domain.User, deactivated bool) ([]*domain.Task, error)
		GetTasksWithItems(u *domain.User, itemLimit int, deactivated bool) ([]*domain.Task, error)
		GetTaskByCode(code string) (*domain.Task, error)
		GetTaskTypeByCode(code string) (*domain.TaskType, error)
		GetTaskCount(active bool) (int, error)

		// items
		CreateItem(*domain.Item, *domain.Task) (*domain.Item, error)
		GetItemCount() (int, error)

		// system summary
		CreateSystemSummary(s *domain.SystemSymmary) (*domain.SystemSymmary, error)
		GetLatestSystemSummary() (*domain.SystemSymmary, error)
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
