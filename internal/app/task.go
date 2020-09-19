package app

import (
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

	_, err = a.repo.CreateTask(t, u, tt)
	if err != nil {
		return err
	}

	return nil
}
