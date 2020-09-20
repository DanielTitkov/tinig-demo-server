// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/systemsummary"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// SystemSummaryDelete is the builder for deleting a SystemSummary entity.
type SystemSummaryDelete struct {
	config
	hooks      []Hook
	mutation   *SystemSummaryMutation
	predicates []predicate.SystemSummary
}

// Where adds a new predicate to the delete builder.
func (ssd *SystemSummaryDelete) Where(ps ...predicate.SystemSummary) *SystemSummaryDelete {
	ssd.predicates = append(ssd.predicates, ps...)
	return ssd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ssd *SystemSummaryDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ssd.hooks) == 0 {
		affected, err = ssd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemSummaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ssd.mutation = mutation
			affected, err = ssd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ssd.hooks) - 1; i >= 0; i-- {
			mut = ssd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ssd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssd *SystemSummaryDelete) ExecX(ctx context.Context) int {
	n, err := ssd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ssd *SystemSummaryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: systemsummary.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemsummary.FieldID,
			},
		},
	}
	if ps := ssd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, ssd.driver, _spec)
}

// SystemSummaryDeleteOne is the builder for deleting a single SystemSummary entity.
type SystemSummaryDeleteOne struct {
	ssd *SystemSummaryDelete
}

// Exec executes the deletion query.
func (ssdo *SystemSummaryDeleteOne) Exec(ctx context.Context) error {
	n, err := ssdo.ssd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{systemsummary.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ssdo *SystemSummaryDeleteOne) ExecX(ctx context.Context) {
	ssdo.ssd.ExecX(ctx)
}