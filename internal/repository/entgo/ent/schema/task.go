package schema

import (
	"regexp"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
	"github.com/facebook/ent/schema/mixin"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").Match(regexp.MustCompile("^[a-zA-Z0-9]+$")),
		field.String("title").NotEmpty(),
		field.String("description").Optional(),
		field.String("code").NotEmpty().Unique().Immutable(),
		field.Bool("active").Default(false),
		field.Time("delete_time").Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		// has
		edge.To("items", Item.Type),
		// belongs to
		edge.From("user", User.Type).Ref("tasks").Unique().Required(),
		edge.From("type", TaskType.Type).Ref("tasks").Unique().Required(),
	}
}

// Indexes of the Task.
func (Task) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("slug").Edges("user").Unique(),
		index.Fields("title").Edges("user").Unique(),
	}
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}
