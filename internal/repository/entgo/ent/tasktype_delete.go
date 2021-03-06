// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TaskTypeDelete is the builder for deleting a TaskType entity.
type TaskTypeDelete struct {
	config
	hooks      []Hook
	mutation   *TaskTypeMutation
	predicates []predicate.TaskType
}

// Where adds a new predicate to the delete builder.
func (ttd *TaskTypeDelete) Where(ps ...predicate.TaskType) *TaskTypeDelete {
	ttd.predicates = append(ttd.predicates, ps...)
	return ttd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ttd *TaskTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ttd.hooks) == 0 {
		affected, err = ttd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttd.mutation = mutation
			affected, err = ttd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ttd.hooks) - 1; i >= 0; i-- {
			mut = ttd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttd *TaskTypeDelete) ExecX(ctx context.Context) int {
	n, err := ttd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ttd *TaskTypeDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: tasktype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tasktype.FieldID,
			},
		},
	}
	if ps := ttd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ttd.driver, _spec)
}

// TaskTypeDeleteOne is the builder for deleting a single TaskType entity.
type TaskTypeDeleteOne struct {
	ttd *TaskTypeDelete
}

// Exec executes the deletion query.
func (ttdo *TaskTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ttdo.ttd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{tasktype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ttdo *TaskTypeDeleteOne) ExecX(ctx context.Context) {
	ttdo.ttd.ExecX(ctx)
}
