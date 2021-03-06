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

// SystemSummaryUpdate is the builder for updating SystemSummary entities.
type SystemSummaryUpdate struct {
	config
	hooks      []Hook
	mutation   *SystemSummaryMutation
	predicates []predicate.SystemSummary
}

// Where adds a new predicate for the builder.
func (ssu *SystemSummaryUpdate) Where(ps ...predicate.SystemSummary) *SystemSummaryUpdate {
	ssu.predicates = append(ssu.predicates, ps...)
	return ssu
}

// SetUsers sets the users field.
func (ssu *SystemSummaryUpdate) SetUsers(i int) *SystemSummaryUpdate {
	ssu.mutation.ResetUsers()
	ssu.mutation.SetUsers(i)
	return ssu
}

// AddUsers adds i to users.
func (ssu *SystemSummaryUpdate) AddUsers(i int) *SystemSummaryUpdate {
	ssu.mutation.AddUsers(i)
	return ssu
}

// SetTasks sets the tasks field.
func (ssu *SystemSummaryUpdate) SetTasks(i int) *SystemSummaryUpdate {
	ssu.mutation.ResetTasks()
	ssu.mutation.SetTasks(i)
	return ssu
}

// AddTasks adds i to tasks.
func (ssu *SystemSummaryUpdate) AddTasks(i int) *SystemSummaryUpdate {
	ssu.mutation.AddTasks(i)
	return ssu
}

// SetActiveTasks sets the active_tasks field.
func (ssu *SystemSummaryUpdate) SetActiveTasks(i int) *SystemSummaryUpdate {
	ssu.mutation.ResetActiveTasks()
	ssu.mutation.SetActiveTasks(i)
	return ssu
}

// AddActiveTasks adds i to active_tasks.
func (ssu *SystemSummaryUpdate) AddActiveTasks(i int) *SystemSummaryUpdate {
	ssu.mutation.AddActiveTasks(i)
	return ssu
}

// SetItems sets the items field.
func (ssu *SystemSummaryUpdate) SetItems(i int) *SystemSummaryUpdate {
	ssu.mutation.ResetItems()
	ssu.mutation.SetItems(i)
	return ssu
}

// AddItems adds i to items.
func (ssu *SystemSummaryUpdate) AddItems(i int) *SystemSummaryUpdate {
	ssu.mutation.AddItems(i)
	return ssu
}

// SetAvgItemsPerTask sets the avg_items_per_task field.
func (ssu *SystemSummaryUpdate) SetAvgItemsPerTask(f float64) *SystemSummaryUpdate {
	ssu.mutation.ResetAvgItemsPerTask()
	ssu.mutation.SetAvgItemsPerTask(f)
	return ssu
}

// AddAvgItemsPerTask adds f to avg_items_per_task.
func (ssu *SystemSummaryUpdate) AddAvgItemsPerTask(f float64) *SystemSummaryUpdate {
	ssu.mutation.AddAvgItemsPerTask(f)
	return ssu
}

// Mutation returns the SystemSummaryMutation object of the builder.
func (ssu *SystemSummaryUpdate) Mutation() *SystemSummaryMutation {
	return ssu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ssu *SystemSummaryUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ssu.hooks) == 0 {
		affected, err = ssu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemSummaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ssu.mutation = mutation
			affected, err = ssu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ssu.hooks) - 1; i >= 0; i-- {
			mut = ssu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ssu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ssu *SystemSummaryUpdate) SaveX(ctx context.Context) int {
	affected, err := ssu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ssu *SystemSummaryUpdate) Exec(ctx context.Context) error {
	_, err := ssu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssu *SystemSummaryUpdate) ExecX(ctx context.Context) {
	if err := ssu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssu *SystemSummaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemsummary.Table,
			Columns: systemsummary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemsummary.FieldID,
			},
		},
	}
	if ps := ssu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ssu.mutation.Users(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldUsers,
		})
	}
	if value, ok := ssu.mutation.AddedUsers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldUsers,
		})
	}
	if value, ok := ssu.mutation.Tasks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldTasks,
		})
	}
	if value, ok := ssu.mutation.AddedTasks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldTasks,
		})
	}
	if value, ok := ssu.mutation.ActiveTasks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldActiveTasks,
		})
	}
	if value, ok := ssu.mutation.AddedActiveTasks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldActiveTasks,
		})
	}
	if value, ok := ssu.mutation.Items(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldItems,
		})
	}
	if value, ok := ssu.mutation.AddedItems(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldItems,
		})
	}
	if value, ok := ssu.mutation.AvgItemsPerTask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemsummary.FieldAvgItemsPerTask,
		})
	}
	if value, ok := ssu.mutation.AddedAvgItemsPerTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemsummary.FieldAvgItemsPerTask,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ssu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemsummary.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// SystemSummaryUpdateOne is the builder for updating a single SystemSummary entity.
type SystemSummaryUpdateOne struct {
	config
	hooks    []Hook
	mutation *SystemSummaryMutation
}

// SetUsers sets the users field.
func (ssuo *SystemSummaryUpdateOne) SetUsers(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.ResetUsers()
	ssuo.mutation.SetUsers(i)
	return ssuo
}

// AddUsers adds i to users.
func (ssuo *SystemSummaryUpdateOne) AddUsers(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.AddUsers(i)
	return ssuo
}

// SetTasks sets the tasks field.
func (ssuo *SystemSummaryUpdateOne) SetTasks(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.ResetTasks()
	ssuo.mutation.SetTasks(i)
	return ssuo
}

// AddTasks adds i to tasks.
func (ssuo *SystemSummaryUpdateOne) AddTasks(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.AddTasks(i)
	return ssuo
}

// SetActiveTasks sets the active_tasks field.
func (ssuo *SystemSummaryUpdateOne) SetActiveTasks(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.ResetActiveTasks()
	ssuo.mutation.SetActiveTasks(i)
	return ssuo
}

// AddActiveTasks adds i to active_tasks.
func (ssuo *SystemSummaryUpdateOne) AddActiveTasks(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.AddActiveTasks(i)
	return ssuo
}

// SetItems sets the items field.
func (ssuo *SystemSummaryUpdateOne) SetItems(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.ResetItems()
	ssuo.mutation.SetItems(i)
	return ssuo
}

// AddItems adds i to items.
func (ssuo *SystemSummaryUpdateOne) AddItems(i int) *SystemSummaryUpdateOne {
	ssuo.mutation.AddItems(i)
	return ssuo
}

// SetAvgItemsPerTask sets the avg_items_per_task field.
func (ssuo *SystemSummaryUpdateOne) SetAvgItemsPerTask(f float64) *SystemSummaryUpdateOne {
	ssuo.mutation.ResetAvgItemsPerTask()
	ssuo.mutation.SetAvgItemsPerTask(f)
	return ssuo
}

// AddAvgItemsPerTask adds f to avg_items_per_task.
func (ssuo *SystemSummaryUpdateOne) AddAvgItemsPerTask(f float64) *SystemSummaryUpdateOne {
	ssuo.mutation.AddAvgItemsPerTask(f)
	return ssuo
}

// Mutation returns the SystemSummaryMutation object of the builder.
func (ssuo *SystemSummaryUpdateOne) Mutation() *SystemSummaryMutation {
	return ssuo.mutation
}

// Save executes the query and returns the updated entity.
func (ssuo *SystemSummaryUpdateOne) Save(ctx context.Context) (*SystemSummary, error) {
	var (
		err  error
		node *SystemSummary
	)
	if len(ssuo.hooks) == 0 {
		node, err = ssuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SystemSummaryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ssuo.mutation = mutation
			node, err = ssuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ssuo.hooks) - 1; i >= 0; i-- {
			mut = ssuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ssuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ssuo *SystemSummaryUpdateOne) SaveX(ctx context.Context) *SystemSummary {
	ss, err := ssuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ss
}

// Exec executes the query on the entity.
func (ssuo *SystemSummaryUpdateOne) Exec(ctx context.Context) error {
	_, err := ssuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ssuo *SystemSummaryUpdateOne) ExecX(ctx context.Context) {
	if err := ssuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ssuo *SystemSummaryUpdateOne) sqlSave(ctx context.Context) (ss *SystemSummary, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   systemsummary.Table,
			Columns: systemsummary.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: systemsummary.FieldID,
			},
		},
	}
	id, ok := ssuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing SystemSummary.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := ssuo.mutation.Users(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldUsers,
		})
	}
	if value, ok := ssuo.mutation.AddedUsers(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldUsers,
		})
	}
	if value, ok := ssuo.mutation.Tasks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldTasks,
		})
	}
	if value, ok := ssuo.mutation.AddedTasks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldTasks,
		})
	}
	if value, ok := ssuo.mutation.ActiveTasks(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldActiveTasks,
		})
	}
	if value, ok := ssuo.mutation.AddedActiveTasks(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldActiveTasks,
		})
	}
	if value, ok := ssuo.mutation.Items(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldItems,
		})
	}
	if value, ok := ssuo.mutation.AddedItems(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: systemsummary.FieldItems,
		})
	}
	if value, ok := ssuo.mutation.AvgItemsPerTask(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemsummary.FieldAvgItemsPerTask,
		})
	}
	if value, ok := ssuo.mutation.AddedAvgItemsPerTask(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: systemsummary.FieldAvgItemsPerTask,
		})
	}
	ss = &SystemSummary{config: ssuo.config}
	_spec.Assign = ss.assignValues
	_spec.ScanValues = ss.scanValues()
	if err = sqlgraph.UpdateNode(ctx, ssuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{systemsummary.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return ss, nil
}
