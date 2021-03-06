// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TaskTypeCreate is the builder for creating a TaskType entity.
type TaskTypeCreate struct {
	config
	mutation *TaskTypeMutation
	hooks    []Hook
}

// SetCode sets the code field.
func (ttc *TaskTypeCreate) SetCode(s string) *TaskTypeCreate {
	ttc.mutation.SetCode(s)
	return ttc
}

// SetTitle sets the title field.
func (ttc *TaskTypeCreate) SetTitle(s string) *TaskTypeCreate {
	ttc.mutation.SetTitle(s)
	return ttc
}

// AddTaskIDs adds the tasks edge to Task by ids.
func (ttc *TaskTypeCreate) AddTaskIDs(ids ...int) *TaskTypeCreate {
	ttc.mutation.AddTaskIDs(ids...)
	return ttc
}

// AddTasks adds the tasks edges to Task.
func (ttc *TaskTypeCreate) AddTasks(t ...*Task) *TaskTypeCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ttc.AddTaskIDs(ids...)
}

// Mutation returns the TaskTypeMutation object of the builder.
func (ttc *TaskTypeCreate) Mutation() *TaskTypeMutation {
	return ttc.mutation
}

// Save creates the TaskType in the database.
func (ttc *TaskTypeCreate) Save(ctx context.Context) (*TaskType, error) {
	if err := ttc.preSave(); err != nil {
		return nil, err
	}
	var (
		err  error
		node *TaskType
	)
	if len(ttc.hooks) == 0 {
		node, err = ttc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ttc.mutation = mutation
			node, err = ttc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ttc.hooks) - 1; i >= 0; i-- {
			mut = ttc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ttc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ttc *TaskTypeCreate) SaveX(ctx context.Context) *TaskType {
	v, err := ttc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ttc *TaskTypeCreate) preSave() error {
	if _, ok := ttc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New("ent: missing required field \"code\"")}
	}
	if v, ok := ttc.mutation.Code(); ok {
		if err := tasktype.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf("ent: validator failed for field \"code\": %w", err)}
		}
	}
	if _, ok := ttc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New("ent: missing required field \"title\"")}
	}
	if v, ok := ttc.mutation.Title(); ok {
		if err := tasktype.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	return nil
}

func (ttc *TaskTypeCreate) sqlSave(ctx context.Context) (*TaskType, error) {
	tt, _spec := ttc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ttc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	tt.ID = int(id)
	return tt, nil
}

func (ttc *TaskTypeCreate) createSpec() (*TaskType, *sqlgraph.CreateSpec) {
	var (
		tt    = &TaskType{config: ttc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tasktype.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tasktype.FieldID,
			},
		}
	)
	if value, ok := ttc.mutation.Code(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasktype.FieldCode,
		})
		tt.Code = value
	}
	if value, ok := ttc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tasktype.FieldTitle,
		})
		tt.Title = value
	}
	if nodes := ttc.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tasktype.TasksTable,
			Columns: []string{tasktype.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return tt, _spec
}

// TaskTypeCreateBulk is the builder for creating a bulk of TaskType entities.
type TaskTypeCreateBulk struct {
	config
	builders []*TaskTypeCreate
}

// Save creates the TaskType entities in the database.
func (ttcb *TaskTypeCreateBulk) Save(ctx context.Context) ([]*TaskType, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ttcb.builders))
	nodes := make([]*TaskType, len(ttcb.builders))
	mutators := make([]Mutator, len(ttcb.builders))
	for i := range ttcb.builders {
		func(i int, root context.Context) {
			builder := ttcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				if err := builder.preSave(); err != nil {
					return nil, err
				}
				mutation, ok := m.(*TaskTypeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ttcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ttcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ttcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX calls Save and panics if Save returns an error.
func (ttcb *TaskTypeCreateBulk) SaveX(ctx context.Context) []*TaskType {
	v, err := ttcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
