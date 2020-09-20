package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// SystemSummary holds the schema definition for the SystemSummary entity.
type SystemSummary struct {
	ent.Schema
}

// Fields of the SystemSummary.
func (SystemSummary) Fields() []ent.Field {
	return []ent.Field{
		field.Int("users"),
		field.Int("tasks"),
		field.Int("active_tasks"),
		field.Int("items"),
		field.Float("avg_items_per_task"),
	}
}

// Edges of the SystemSummary.
func (SystemSummary) Edges() []ent.Edge {
	return nil
}

func (SystemSummary) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.CreateTime{},
	}
}
