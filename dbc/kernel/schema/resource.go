package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Status ...
// ENUM(none,waiting,success,failed,max)
type Status uint8

// Step ...
// ENUM(none,add,remove,max)
type Step uint8

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("rid").Unique(), //pin hash
		field.Uint8("status").Default(0),
		field.Uint8("step").Default(0),
		field.Int("retried").Default(0),
		field.Int("priority").Default(-1), //priority(-1:never called)
		field.String("relate").Default("none"),
		field.Int64("updated_unix").Default(0),
		field.String("comment").Optional(), //failed log
	}
}

func (Resource) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("rid"),
	}
}
