package app

import (
	"errors"
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

func (a *App) GetTasks(u *domain.User, withItems bool, deactivated bool) ([]*domain.Task, error) {
	u, err := a.repo.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	var itemLimit int
	if withItems {
		itemLimit = a.cfg.Task.ItemLimit
	}

	return a.repo.GetTasksWithItems(u, itemLimit, deactivated)
}

func (a *App) UpdateTask(t *domain.Task) error {
	err := a.ValidateTaskUser(t)
	if err != nil {
		return err
	}

	err = a.ValidateTaskParams(t)
	if err != nil {
		return err
	}

	_, err = a.repo.UpdateTask(t)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) ValidateTaskUser(t *domain.Task) error {
	user, err := a.repo.GetUserByUsername(t.User)
	if err != nil {
		return err
	}

	task, err := a.repo.GetTaskByCode(t.Code)
	if err != nil {
		return err
	}

	if task.User != user.Username {
		return errors.New("trying to update task that belongs to another user")
	}

	return nil
}

func (a *App) ValidateTaskParams(t *domain.Task) error {
	task, err := a.repo.GetTaskByCode(t.Code)
	if err != nil {
		return err
	}

	taskType := task.Type
	params := t.Params
	switch taskType {
	case domain.TaskTypeRandom: // TODO: move params to config
		if min := params.Random.Min; min == 0 || min < 1 || min > 100 {
			return errors.New("min values are allowed in [1, 99] for task type " + taskType)
		}
		if max := params.Random.Max; max == 0 || max < 2 || max > 101 {
			return errors.New("max values are allowed in [2, 100] for task type " + taskType)
		}
		if params.Random.Min > params.Random.Max {
			return errors.New("min can not be bigger than max")
		}
	case domain.TaskTypePrice:
		return errors.New("price task params are not implemented yet")
	default:
		return errors.New("got unsupported task type: " + taskType)
	}

	return nil
}
