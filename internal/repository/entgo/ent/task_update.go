// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/domain"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/item"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/task"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/tasktype"
	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/user"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks      []Hook
	mutation   *TaskMutation
	predicates []predicate.Task
}

// Where adds a new predicate for the builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.predicates = append(tu.predicates, ps...)
	return tu
}

// SetSlug sets the slug field.
func (tu *TaskUpdate) SetSlug(s string) *TaskUpdate {
	tu.mutation.SetSlug(s)
	return tu
}

// SetTitle sets the title field.
func (tu *TaskUpdate) SetTitle(s string) *TaskUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetDescription sets the description field.
func (tu *TaskUpdate) SetDescription(s string) *TaskUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the description field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDescription(s *string) *TaskUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// ClearDescription clears the value of description.
func (tu *TaskUpdate) ClearDescription() *TaskUpdate {
	tu.mutation.ClearDescription()
	return tu
}

// SetActive sets the active field.
func (tu *TaskUpdate) SetActive(b bool) *TaskUpdate {
	tu.mutation.SetActive(b)
	return tu
}

// SetNillableActive sets the active field if the given value is not nil.
func (tu *TaskUpdate) SetNillableActive(b *bool) *TaskUpdate {
	if b != nil {
		tu.SetActive(*b)
	}
	return tu
}

// SetDeleteTime sets the delete_time field.
func (tu *TaskUpdate) SetDeleteTime(t time.Time) *TaskUpdate {
	tu.mutation.SetDeleteTime(t)
	return tu
}

// SetNillableDeleteTime sets the delete_time field if the given value is not nil.
func (tu *TaskUpdate) SetNillableDeleteTime(t *time.Time) *TaskUpdate {
	if t != nil {
		tu.SetDeleteTime(*t)
	}
	return tu
}

// ClearDeleteTime clears the value of delete_time.
func (tu *TaskUpdate) ClearDeleteTime() *TaskUpdate {
	tu.mutation.ClearDeleteTime()
	return tu
}

// SetMeta sets the meta field.
func (tu *TaskUpdate) SetMeta(dm domain.TaskMeta) *TaskUpdate {
	tu.mutation.SetMeta(dm)
	return tu
}

// SetNillableMeta sets the meta field if the given value is not nil.
func (tu *TaskUpdate) SetNillableMeta(dm *domain.TaskMeta) *TaskUpdate {
	if dm != nil {
		tu.SetMeta(*dm)
	}
	return tu
}

// ClearMeta clears the value of meta.
func (tu *TaskUpdate) ClearMeta() *TaskUpdate {
	tu.mutation.ClearMeta()
	return tu
}

// AddItemIDs adds the items edge to Item by ids.
func (tu *TaskUpdate) AddItemIDs(ids ...int) *TaskUpdate {
	tu.mutation.AddItemIDs(ids...)
	return tu
}

// AddItems adds the items edges to Item.
func (tu *TaskUpdate) AddItems(i ...*Item) *TaskUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tu.AddItemIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (tu *TaskUpdate) SetUserID(id int) *TaskUpdate {
	tu.mutation.SetUserID(id)
	return tu
}

// SetUser sets the user edge to User.
func (tu *TaskUpdate) SetUser(u *User) *TaskUpdate {
	return tu.SetUserID(u.ID)
}

// SetTypeID sets the type edge to TaskType by id.
func (tu *TaskUpdate) SetTypeID(id int) *TaskUpdate {
	tu.mutation.SetTypeID(id)
	return tu
}

// SetType sets the type edge to TaskType.
func (tu *TaskUpdate) SetType(t *TaskType) *TaskUpdate {
	return tu.SetTypeID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// RemoveItemIDs removes the items edge to Item by ids.
func (tu *TaskUpdate) RemoveItemIDs(ids ...int) *TaskUpdate {
	tu.mutation.RemoveItemIDs(ids...)
	return tu
}

// RemoveItems removes items edges to Item.
func (tu *TaskUpdate) RemoveItems(i ...*Item) *TaskUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tu.RemoveItemIDs(ids...)
}

// ClearUser clears the user edge to User.
func (tu *TaskUpdate) ClearUser() *TaskUpdate {
	tu.mutation.ClearUser()
	return tu
}

// ClearType clears the type edge to TaskType.
func (tu *TaskUpdate) ClearType() *TaskUpdate {
	tu.mutation.ClearType()
	return tu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := task.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
	if v, ok := tu.mutation.Slug(); ok {
		if err := task.SlugValidator(v); err != nil {
			return 0, &ValidationError{Name: "slug", err: fmt.Errorf("ent: validator failed for field \"slug\": %w", err)}
		}
	}
	if v, ok := tu.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return 0, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}

	if _, ok := tu.mutation.UserID(); tu.mutation.UserCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"user\"")
	}

	if _, ok := tu.mutation.TypeID(); tu.mutation.TypeCleared() && !ok {
		return 0, errors.New("ent: clearing a unique edge \"type\"")
	}
	var (
		err      error
		affected int
	)
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	if ps := tu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldSlug,
		})
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tu.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldActive,
		})
	}
	if value, ok := tu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDeleteTime,
		})
	}
	if tu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDeleteTime,
		})
	}
	if tu.mutation.ParamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldParams,
		})
	}
	if value, ok := tu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: task.FieldMeta,
		})
	}
	if tu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldMeta,
		})
	}
	if nodes := tu.mutation.RemovedItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ItemsTable,
			Columns: []string{task.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ItemsTable,
			Columns: []string{task.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TypeTable,
			Columns: []string{task.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TypeTable,
			Columns: []string{task.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// SetSlug sets the slug field.
func (tuo *TaskUpdateOne) SetSlug(s string) *TaskUpdateOne {
	tuo.mutation.SetSlug(s)
	return tuo
}

// SetTitle sets the title field.
func (tuo *TaskUpdateOne) SetTitle(s string) *TaskUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetDescription sets the description field.
func (tuo *TaskUpdateOne) SetDescription(s string) *TaskUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the description field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDescription(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// ClearDescription clears the value of description.
func (tuo *TaskUpdateOne) ClearDescription() *TaskUpdateOne {
	tuo.mutation.ClearDescription()
	return tuo
}

// SetActive sets the active field.
func (tuo *TaskUpdateOne) SetActive(b bool) *TaskUpdateOne {
	tuo.mutation.SetActive(b)
	return tuo
}

// SetNillableActive sets the active field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableActive(b *bool) *TaskUpdateOne {
	if b != nil {
		tuo.SetActive(*b)
	}
	return tuo
}

// SetDeleteTime sets the delete_time field.
func (tuo *TaskUpdateOne) SetDeleteTime(t time.Time) *TaskUpdateOne {
	tuo.mutation.SetDeleteTime(t)
	return tuo
}

// SetNillableDeleteTime sets the delete_time field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableDeleteTime(t *time.Time) *TaskUpdateOne {
	if t != nil {
		tuo.SetDeleteTime(*t)
	}
	return tuo
}

// ClearDeleteTime clears the value of delete_time.
func (tuo *TaskUpdateOne) ClearDeleteTime() *TaskUpdateOne {
	tuo.mutation.ClearDeleteTime()
	return tuo
}

// SetMeta sets the meta field.
func (tuo *TaskUpdateOne) SetMeta(dm domain.TaskMeta) *TaskUpdateOne {
	tuo.mutation.SetMeta(dm)
	return tuo
}

// SetNillableMeta sets the meta field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableMeta(dm *domain.TaskMeta) *TaskUpdateOne {
	if dm != nil {
		tuo.SetMeta(*dm)
	}
	return tuo
}

// ClearMeta clears the value of meta.
func (tuo *TaskUpdateOne) ClearMeta() *TaskUpdateOne {
	tuo.mutation.ClearMeta()
	return tuo
}

// AddItemIDs adds the items edge to Item by ids.
func (tuo *TaskUpdateOne) AddItemIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.AddItemIDs(ids...)
	return tuo
}

// AddItems adds the items edges to Item.
func (tuo *TaskUpdateOne) AddItems(i ...*Item) *TaskUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tuo.AddItemIDs(ids...)
}

// SetUserID sets the user edge to User by id.
func (tuo *TaskUpdateOne) SetUserID(id int) *TaskUpdateOne {
	tuo.mutation.SetUserID(id)
	return tuo
}

// SetUser sets the user edge to User.
func (tuo *TaskUpdateOne) SetUser(u *User) *TaskUpdateOne {
	return tuo.SetUserID(u.ID)
}

// SetTypeID sets the type edge to TaskType by id.
func (tuo *TaskUpdateOne) SetTypeID(id int) *TaskUpdateOne {
	tuo.mutation.SetTypeID(id)
	return tuo
}

// SetType sets the type edge to TaskType.
func (tuo *TaskUpdateOne) SetType(t *TaskType) *TaskUpdateOne {
	return tuo.SetTypeID(t.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// RemoveItemIDs removes the items edge to Item by ids.
func (tuo *TaskUpdateOne) RemoveItemIDs(ids ...int) *TaskUpdateOne {
	tuo.mutation.RemoveItemIDs(ids...)
	return tuo
}

// RemoveItems removes items edges to Item.
func (tuo *TaskUpdateOne) RemoveItems(i ...*Item) *TaskUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return tuo.RemoveItemIDs(ids...)
}

// ClearUser clears the user edge to User.
func (tuo *TaskUpdateOne) ClearUser() *TaskUpdateOne {
	tuo.mutation.ClearUser()
	return tuo
}

// ClearType clears the type edge to TaskType.
func (tuo *TaskUpdateOne) ClearType() *TaskUpdateOne {
	tuo.mutation.ClearType()
	return tuo
}

// Save executes the query and returns the updated entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := task.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
	if v, ok := tuo.mutation.Slug(); ok {
		if err := task.SlugValidator(v); err != nil {
			return nil, &ValidationError{Name: "slug", err: fmt.Errorf("ent: validator failed for field \"slug\": %w", err)}
		}
	}
	if v, ok := tuo.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return nil, &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}

	if _, ok := tuo.mutation.UserID(); tuo.mutation.UserCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"user\"")
	}

	if _, ok := tuo.mutation.TypeID(); tuo.mutation.TypeCleared() && !ok {
		return nil, errors.New("ent: clearing a unique edge \"type\"")
	}
	var (
		err  error
		node *Task
	)
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	t, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return t
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (t *Task, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   task.Table,
			Columns: task.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: task.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Task.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.Slug(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldSlug,
		})
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: task.FieldDescription,
		})
	}
	if tuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: task.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.Active(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: task.FieldActive,
		})
	}
	if value, ok := tuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: task.FieldDeleteTime,
		})
	}
	if tuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: task.FieldDeleteTime,
		})
	}
	if tuo.mutation.ParamsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldParams,
		})
	}
	if value, ok := tuo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: task.FieldMeta,
		})
	}
	if tuo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: task.FieldMeta,
		})
	}
	if nodes := tuo.mutation.RemovedItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ItemsTable,
			Columns: []string{task.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.ItemsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   task.ItemsTable,
			Columns: []string{task.ItemsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.UserTable,
			Columns: []string{task.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TypeTable,
			Columns: []string{task.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.TypeTable,
			Columns: []string{task.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tasktype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	t = &Task{config: tuo.config}
	_spec.Assign = t.assignValues
	_spec.ScanValues = t.scanValues()
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{task.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return t, nil
}
