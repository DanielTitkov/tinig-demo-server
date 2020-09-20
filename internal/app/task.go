package app

import (
	"strconv"
	"strings"
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

func (a *App) CreateTask(t *domain.Task) error {
	u, err := a.repo.GetUserByUsername(t.User)
	if err != nil {
		return err
	}

	tt, err := a.repo.GetTaskTypeByCode(t.Type)
	if err != nil {
		return err
	}

	ts := strconv.FormatInt(time.Now().Unix(), 10)
	code := strings.Join([]string{t.User, t.Type, t.Slug, ts}, "_")
	t.Code = code
	_, err = a.repo.CreateTask(t, u, tt)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) GetTasks(u *domain.User) ([]*domain.Task, error) {
	u, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	return a.repo.GetTasks(u)
}

func (a *App) GetTasksWithItems(u *domain.User) ([]*domain.TaskWithItems, error) {
	u, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	return a.repo.GetTasksWithItems(u, a.cfg.Task.ItemLimit)
}
