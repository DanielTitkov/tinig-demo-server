package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
)

func (r *EntgoRepository) CreateTask(t *domain.Task, u *domain.User, tt *domain.TaskType) (*domain.Task, error) {
	task, err := r.client.Task.
		Create().
		SetTitle(t.Title).
		SetUserID(u.ID). // let's see what if type is not set
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:    task.ID,
		Title: t.Title,
		User:  t.User,
		Type:  t.Type,
	}, nil
}

func (r *EntgoRepository) GetTaskTypeByCode(code string) (*domain.TaskType, error) {
	taskType, err := r.client.TaskType.
		Query().
		Where(tasktype.CodeEQ(code)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.TaskType{
		ID:   taskType.ID,
		Code: taskType.Code,
	}, nil
}
