package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Instruct holds the schema definition for the Instruct entity.
type Instruct struct {
	ent.Schema
}

// Fields of the Version.
func (Instruct) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("current_unix").Default(0),
		field.Int64("updated_unix").Default(0),
	}
}

// Edges of the Version.
func (Instruct) Edges() []ent.Edge {
	return nil
}
