package entgo

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/user"
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

func (r *EntgoRepository) GetTasksWithItems(u *domain.User, itemLimit int) ([]*domain.TaskWithItems, error) {
	tasks, err := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryTasks().
		WithType().
		WithItems(func(q *ent.ItemQuery) {
			q.Limit(1)
		}).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	fmt.Printf("%+v", tasks)

	var res []*domain.TaskWithItems
	for _, t := range tasks {
		var items []domain.Item
		fmt.Printf("%+v", t.Edges.Items)
		for _, i := range t.Edges.Items {
			if i.Data == nil {
				continue // items without data don't matter
			}
			items = append(items, domain.Item{
				ID:     i.ID,
				Hash:   i.Hash,
				Source: i.Source,
				Data:   *i.Data,
				Task:   t.Code,
			})
		}

		res = append(res, &domain.TaskWithItems{
			Task: domain.Task{ // TODO: maybe use pointer
				ID:          t.ID,
				Code:        t.Code,
				Slug:        t.Slug,
				Title:       t.Title,
				Description: t.Description,
				Type:        t.Edges.Type.Code,
				User:        u.Username,
			},
			Items: items,
		})
	}

	return res, nil
}

func (r *EntgoRepository) GetTasks(u *domain.User) ([]*domain.Task, error) {
	tasks, err := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryTasks().
		WithType().
		All(context.Background())
	if err != nil {
		return nil, err
	}

	var res []*domain.Task
	for _, t := range tasks {
		res = append(res, &domain.Task{
			ID:          t.ID,
			Code:        t.Code,
			Slug:        t.Slug,
			Title:       t.Title,
			Description: t.Description,
			Type:        t.Edges.Type.Code,
			User:        u.Username,
		})
	}

	return res, nil
}
