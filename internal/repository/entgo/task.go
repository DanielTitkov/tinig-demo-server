package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/user"
	"github.com/bradfitz/slice"
)

func (r *EntgoRepository) CreateTask(t *domain.Task, u *domain.User, tt *domain.TaskType) (*domain.Task, error) {
	task, err := r.client.Task.
		Create().
		SetSlug(t.Slug).
		SetCode(t.Code).
		SetTitle(t.Title).
		SetDescription(t.Description).
		SetParams(t.Params).
		SetMeta(t.Meta).
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
		Params:      t.Params,
		Meta:        t.Meta,
	}, nil
}

func (r *EntgoRepository) GetTaskByCode(code string) (*domain.Task, error) {
	task, err := r.client.Task.
		Query().
		Where(task.CodeEQ(code)).
		WithUser().
		WithType().
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
		Params:      task.Params,
		Meta:        task.Meta,
		User:        task.Edges.User.Username,
		Type:        task.Edges.Type.Code,
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

func (r *EntgoRepository) UpdateTask(t *domain.Task) (*domain.Task, error) {
	task, err := r.client.Task.
		Query().
		Where(task.CodeEQ(t.Code)).
		Only(context.Background())
	if err != nil {
		return nil, err
	}

	title := task.Title
	if t.Title != "" {
		title = t.Title
	}

	description := task.Description
	if t.Description != "" {
		description = t.Description
	}

	// TODO: maybe create separate methods for different fields

	task, err = task.Update().
		SetTitle(title).
		SetDescription(description).
		SetActive(t.Active). // required
		SetMeta(t.Meta).     // TODO: maybe this should have special logic
		Save(context.Background())

	if err != nil {
		return nil, err
	}

	return &domain.Task{
		ID:          task.ID,
		Code:        task.Code,
		Title:       task.Title,
		Slug:        task.Slug,
		Description: task.Description,
		Params:      task.Params,
		Meta:        task.Meta,
		// User:     task.Edges.User.Username,
		// Type:     task.Edges.Type.Code,
	}, nil
}

func (r *EntgoRepository) GetTasks(u *domain.User) ([]*domain.Task, error) {
	tasks, err := r.client.User.
		Query().
		Where(user.UsernameEQ(u.Username)).
		QueryTasks().
		Where(task.DeleteTimeIsNil()).
		WithType().
		Order(ent.Desc(task.FieldUpdateTime)).
		All(context.Background())
	if err != nil {
		return nil, err
	}

	var res []*domain.Task
	for _, t := range tasks {
		res = append(res, &domain.Task{ // TODO: maybe use pointer
			ID:          t.ID,
			Code:        t.Code,
			Slug:        t.Slug,
			Title:       t.Title,
			Description: t.Description,
			Active:      t.Active,
			Params:      t.Params,
			Meta:        t.Meta,
			Type:        t.Edges.Type.Code,
			User:        u.Username,
		})
	}

	return res, nil
}

func (r *EntgoRepository) GetTasksWithItems(u *domain.User, itemLimit int) ([]*domain.Task, error) {
	tasks, err := r.GetTasks(u) // FIXME: maybe there is the way to do in one query? Ask A.R.
	if err != nil {
		return nil, err
	}

	if itemLimit == 0 {
		return tasks, nil
	}

	var res []*domain.Task
	for _, fetchedTask := range tasks {
		t, err := r.client.Task.
			Query().
			Where(task.IDEQ(fetchedTask.ID)).
			WithItems(func(q *ent.ItemQuery) {
				q.Order(ent.Desc(item.FieldCreateTime))
				q.Limit(itemLimit)
			}).
			Only(context.Background())
		if err != nil {
			return nil, err
		}
		var items []domain.Item
		for _, i := range t.Edges.Items {
			items = append(items, domain.Item{
				ID:         i.ID,
				Hash:       i.Hash,
				Source:     i.Source,
				Data:       i.Data,
				Task:       t.Code,
				CreateTime: i.CreateTime,
			})
		}

		slice.Sort(items, func(i, j int) bool {
			return items[j].CreateTime.After(items[i].CreateTime)
		})

		res = append(res, &domain.Task{ // TODO: maybe use pointer
			ID:          t.ID,
			Code:        t.Code,
			Slug:        t.Slug,
			Title:       t.Title,
			Description: t.Description,
			Active:      t.Active,
			Params:      t.Params,
			Meta:        t.Meta,
			Type:        fetchedTask.Type,
			User:        u.Username,
			Items:       items,
		})
	}

	return res, nil
}

func (r *EntgoRepository) GetTaskCount(active bool) (int, error) {
	taskCountQuery := r.client.Task.Query()
	if active {
		taskCountQuery = taskCountQuery.Where(task.ActiveEQ(true))
	}

	return taskCountQuery.Count(context.Background())
}
