package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// TaskType holds the schema definition for the TaskType entity.
type TaskType struct {
	ent.Schema
}

// Fields of the TaskType.
func (TaskType) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").Unique().NotEmpty(),
		field.String("title").Unique().NotEmpty(),
	}
}

// Edges of the TaskType.
func (TaskType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
	}
}
