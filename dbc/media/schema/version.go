package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Version holds the schema definition for the Version entity.
type Version struct {
	ent.Schema
}

const CurrentVersion = 1
const CurrentDataVersion = 1

// Fields of the Version.
func (Version) Fields() []ent.Field {
	return []ent.Field{
		field.Int("Current").Positive().Default(CurrentVersion),
		field.Int("Last").Positive().Default(CurrentVersion),
	}
}

// Edges of the Version.
func (Version) Edges() []ent.Edge {
	return nil
}
