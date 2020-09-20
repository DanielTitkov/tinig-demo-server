// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/DanielTitkov/tinig-demo-server/internal/repository/entgo/ent/systemsummary"
	"github.com/facebook/ent/dialect/sql"
)

// SystemSummary is the model entity for the SystemSummary schema.
type SystemSummary struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// Users holds the value of the "users" field.
	Users int `json:"users,omitempty"`
	// Tasks holds the value of the "tasks" field.
	Tasks int `json:"tasks,omitempty"`
	// ActiveTasks holds the value of the "active_tasks" field.
	ActiveTasks int `json:"active_tasks,omitempty"`
	// Items holds the value of the "items" field.
	Items int `json:"items,omitempty"`
	// AvgItemsPerTask holds the value of the "avg_items_per_task" field.
	AvgItemsPerTask float64 `json:"avg_items_per_task,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*SystemSummary) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},   // id
		&sql.NullTime{},    // create_time
		&sql.NullInt64{},   // users
		&sql.NullInt64{},   // tasks
		&sql.NullInt64{},   // active_tasks
		&sql.NullInt64{},   // items
		&sql.NullFloat64{}, // avg_items_per_task
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the SystemSummary fields.
func (ss *SystemSummary) assignValues(values ...interface{}) error {
	if m, n := len(values), len(systemsummary.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	ss.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field create_time", values[0])
	} else if value.Valid {
		ss.CreateTime = value.Time
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field users", values[1])
	} else if value.Valid {
		ss.Users = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field tasks", values[2])
	} else if value.Valid {
		ss.Tasks = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field active_tasks", values[3])
	} else if value.Valid {
		ss.ActiveTasks = int(value.Int64)
	}
	if value, ok := values[4].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field items", values[4])
	} else if value.Valid {
		ss.Items = int(value.Int64)
	}
	if value, ok := values[5].(*sql.NullFloat64); !ok {
		return fmt.Errorf("unexpected type %T for field avg_items_per_task", values[5])
	} else if value.Valid {
		ss.AvgItemsPerTask = value.Float64
	}
	return nil
}

// Update returns a builder for updating this SystemSummary.
// Note that, you need to call SystemSummary.Unwrap() before calling this method, if this SystemSummary
// was returned from a transaction, and the transaction was committed or rolled back.
func (ss *SystemSummary) Update() *SystemSummaryUpdateOne {
	return (&SystemSummaryClient{config: ss.config}).UpdateOne(ss)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ss *SystemSummary) Unwrap() *SystemSummary {
	tx, ok := ss.config.driver.(*txDriver)
	if !ok {
		panic("ent: SystemSummary is not a transactional entity")
	}
	ss.config.driver = tx.drv
	return ss
}

// String implements the fmt.Stringer.
func (ss *SystemSummary) String() string {
	var builder strings.Builder
	builder.WriteString("SystemSummary(")
	builder.WriteString(fmt.Sprintf("id=%v", ss.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(ss.CreateTime.Format(time.ANSIC))
	builder.WriteString(", users=")
	builder.WriteString(fmt.Sprintf("%v", ss.Users))
	builder.WriteString(", tasks=")
	builder.WriteString(fmt.Sprintf("%v", ss.Tasks))
	builder.WriteString(", active_tasks=")
	builder.WriteString(fmt.Sprintf("%v", ss.ActiveTasks))
	builder.WriteString(", items=")
	builder.WriteString(fmt.Sprintf("%v", ss.Items))
	builder.WriteString(", avg_items_per_task=")
	builder.WriteString(fmt.Sprintf("%v", ss.AvgItemsPerTask))
	builder.WriteByte(')')
	return builder.String()
}

// SystemSummaries is a parsable slice of SystemSummary.
type SystemSummaries []*SystemSummary

func (ss SystemSummaries) config(cfg config) {
	for _i := range ss {
		ss[_i].config = cfg
	}
}