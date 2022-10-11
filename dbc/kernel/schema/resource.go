package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Status ...
// ENUM(none,waiting,success,failed,max)
type Status uint32

// Step ...
// ENUM(none,add,remove,max)
type Step uint32

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("rid").Unique(), //pin hash
		field.Uint32("status").Default(0),
		field.Int("retries").Default(0),
		field.Uint32("step").Default(0),
		field.Int("priority").Default(0), //优先级
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
