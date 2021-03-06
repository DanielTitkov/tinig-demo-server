// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks      []Hook
	mutation   *ItemMutation
	predicates []predicate.Item
}

// Where adds a new predicate for the builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// SetSource sets the source field.
func (iu *ItemUpdate) SetSource(s string) *ItemUpdate {
	iu.mutation.SetSource(s)
	return iu
}

// SetHash sets the hash field.
func (iu *ItemUpdate) SetHash(s string) *ItemUpdate {
	iu.mutation.SetHash(s)
	return iu
}

// SetData sets the data field.
func (iu *ItemUpdate) SetData(dd domain.ItemData) *ItemUpdate {
	iu.mutation.SetData(dd)
	return iu
}

// SetNillableData sets the data field if the given value is not nil.
func (iu *ItemUpdate) SetNillableData(dd *domain.ItemData) *ItemUpdate {
	if dd != nil {
		iu.SetData(*dd)
	}
	return iu
}

// ClearData clears the value of data.
func (iu *ItemUpdate) ClearData() *ItemUpdate {
	iu.mutation.ClearData()
	return iu
}

// SetTaskID sets the task edge to Task by id.
func (iu *ItemUpdate) SetTaskID(id int) *ItemUpdate {
	iu.mutation.SetTaskID(id)
	return iu
}

// SetTask sets the task edge to Task.
func (iu *ItemUpdate) SetTask(t *Task) *ItemUpdate {
	return iu.SetTaskID(t.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearTask clears the task edge to Task.
func (iu *ItemUpdate) ClearTask() *ItemUpdate {
	iu.mutation.ClearTask()
	return iu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
	if v, ok := iu.mutation.Source(); ok {
		if err := item.SourceValidator(v); err != nil {
			return 0, &ValidationError{Name: "source", err: fmt.Errorf("ent: validator failed for field \"source\": %w", err)}
		}
	}
	if v, ok := iu.mutation.Hash(); ok {
		if err := item.HashValidator(v); err != nil {
			return 0, &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}

	if _, ok := iu.mutation.TaskID(); iu.mutation.TaskCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"task\"")
	}
	var (
		err      error
		affected int
	)
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldSource,
		})
	}
	if value, ok := iu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldHash,
		})
	}
	if value, ok := iu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: item.FieldData,
		})
	}
	if iu.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: item.FieldData,
		})
	}
	if iu.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskTable,
			Columns: []string{item.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskTable,
			Columns: []string{item.TaskColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// SetSource sets the source field.
func (iuo *ItemUpdateOne) SetSource(s string) *ItemUpdateOne {
	iuo.mutation.SetSource(s)
	return iuo
}

// SetHash sets the hash field.
func (iuo *ItemUpdateOne) SetHash(s string) *ItemUpdateOne {
	iuo.mutation.SetHash(s)
	return iuo
}

// SetData sets the data field.
func (iuo *ItemUpdateOne) SetData(dd domain.ItemData) *ItemUpdateOne {
	iuo.mutation.SetData(dd)
	return iuo
}

// SetNillableData sets the data field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableData(dd *domain.ItemData) *ItemUpdateOne {
	if dd != nil {
		iuo.SetData(*dd)
	}
	return iuo
}

// ClearData clears the value of data.
func (iuo *ItemUpdateOne) ClearData() *ItemUpdateOne {
	iuo.mutation.ClearData()
	return iuo
}

// SetTaskID sets the task edge to Task by id.
func (iuo *ItemUpdateOne) SetTaskID(id int) *ItemUpdateOne {
	iuo.mutation.SetTaskID(id)
	return iuo
}

// SetTask sets the task edge to Task.
func (iuo *ItemUpdateOne) SetTask(t *Task) *ItemUpdateOne {
	return iuo.SetTaskID(t.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearTask clears the task edge to Task.
func (iuo *ItemUpdateOne) ClearTask() *ItemUpdateOne {
	iuo.mutation.ClearTask()
	return iuo
}

// Save executes the query and returns the updated entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
	if v, ok := iuo.mutation.Source(); ok {
		if err := item.SourceValidator(v); err != nil {
			return nil, &ValidationError{Name: "source", err: fmt.Errorf("ent: validator failed for field \"source\": %w", err)}
		}
	}
	if v, ok := iuo.mutation.Hash(); ok {
		if err := item.HashValidator(v); err != nil {
			return nil, &ValidationError{Name: "hash", err: fmt.Errorf("ent: validator failed for field \"hash\": %w", err)}
		}
	}

	if _, ok := iuo.mutation.TaskID(); iuo.mutation.TaskCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"task\"")
	}
	var (
		err  error
		node *Item
	)
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (i *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Item.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldSource,
		})
	}
	if value, ok := iuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldHash,
		})
	}
	if value, ok := iuo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: item.FieldData,
		})
	}
	if iuo.mutation.DataCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: item.FieldData,
		})
	}
	if iuo.mutation.TaskCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskTable,
			Columns: []string{item.TaskColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.TaskIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.TaskTable,
			Columns: []string{item.TaskColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	i = &Item{config: iuo.config}
	_spec.Assign = i.assignValues
	_spec.ScanValues = i.scanValues()
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return i, nil
}
