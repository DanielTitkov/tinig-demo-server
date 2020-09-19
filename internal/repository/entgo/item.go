package entgo

import (
	"context"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
)

func (r *EntgoRepository) CreateItem(i *domain.Item, t *domain.Task) (*domain.Item, error) {
	item, err := r.client.Item.
		Create().
		SetSource(i.Source).
		SetHash(i.Hash).
		SetData(&i.Data).
		SetTaskID(t.ID).
		Save(context.Background())
	if err != nil {
		return nil, err
	}

	return &domain.Item{
		ID:     item.ID,
		Source: item.Source,
		Hash:   item.Hash,
		Task:   i.Task,
	}, nil
}
