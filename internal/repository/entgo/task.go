package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/item"
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

func (r *EntgoRepository) GetTasks(u *domain.User, itemLimit int, deactivated bool) ([]*domain.Task, error) {
	taskBaseQuery := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryTasks().
		Where(task.ActiveEQ(!deactivated)).
		WithType()
	taskQuery := taskBaseQuery
	if itemLimit > 0 {
		taskQuery = taskBaseQuery.
			WithItems(func(q *ent.ItemQuery) {
				q.Order(ent.Desc(item.FieldCreateTime))
				q.Limit(itemLimit)
			})
	}
	tasks, err := taskQuery.All(context.Background())
	if err != nil {
		return nil, err
	}

	var res []*domain.Task
	for _, t := range tasks {
		var items []domain.Item
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

		res = append(res, &domain.Task{ // TODO: maybe use pointer
			ID:          t.ID,
			Code:        t.Code,
			Slug:        t.Slug,
			Title:       t.Title,
			Description: t.Description,
			Type:        t.Edges.Type.Code,
			Active:      t.Active,
			User:        u.Username,
			Items:       items,
		})
	}

	return res, nil
}
