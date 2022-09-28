package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

const (
	PinWaiting  = "waiting"
	PinPinning  = "pinning"
	PinSuccess  = "success"
	PinFailed   = "failed"
	PinNotFound = "notfound"
)

const (
	StepNone   = "none"
	StepAdd    = "add"
	StepRemove = "remove"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("rid").Unique(), //pin hash
		field.String("status").Default("waiting"),
		field.String("step").Default("none"),
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
