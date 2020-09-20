package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/systemsummary"
)

func (r *EntgoRepository) CreateSystemSummary(s *domain.SystemSymmary) (*domain.SystemSymmary, error) {
	sum, err := r.client.SystemSummary.
		Create().
		SetUsers(s.Users).
		SetTasks(s.Tasks).
		SetActiveTasks(s.ActiveTasks).
		SetAvgItemsPerTask(s.AvgItemsPerTask).
		SetItems(s.Items).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.SystemSymmary{
		ID:              sum.ID,
		CreateTime:      sum.CreateTime,
		Tasks:           sum.Tasks,
		ActiveTasks:     sum.ActiveTasks,
		Users:           sum.Users,
		AvgItemsPerTask: sum.AvgItemsPerTask,
		Items:           sum.Items,
	}, nil
}

func (r *EntgoRepository) GetLatestSystemSummary() (*domain.SystemSymmary, error) {
	sum, err := r.client.SystemSummary.
		Query().
		Order(ent.Desc(systemsummary.FieldCreateTime)).
		Limit(1).
		First(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.SystemSymmary{
		ID:              sum.ID,
		Users:           sum.Users,
		Tasks:           sum.Tasks,
		ActiveTasks:     sum.ActiveTasks,
		Items:           sum.Items,
		AvgItemsPerTask: sum.AvgItemsPerTask,
		CreateTime:      sum.CreateTime,
	}, nil
}
