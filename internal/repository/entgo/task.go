package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
)

func (r *EntgoRepository) CreateTask(t *domain.Task, u *domain.User, tt *domain.TaskType) (*domain.Task, error) {
	task, err := r.client.Task.
		Create().
		SetSlug(t.Slug).
		SetCode(t.Code).
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetUserID(u.ID).
		SetTypeID(tt.ID).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          task.ID,
		User:        t.User,
		Type:        t.Type,
		Slug:        t.Slug,
		Code:        t.Code,
		Title:       t.Title,
		Description: t.Description,
	}, nil
}

func (r *EntgoRepository) GetTaskByCode(code string) (*domain.Task, error) {
	task, err := r.client.Task.
		Query().
		Where(task.CodeEQ(code)).
		WithUser().
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          task.ID,
		Code:        task.Code,
		Title:       task.Title,
		Slug:        task.Slug,
		Description: task.Description,
		User:        task.Edges.User.Username,
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
