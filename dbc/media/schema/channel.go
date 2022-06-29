package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/tikafog/of/dbc/mixin"
)

// Channel holds the schema definition for the Channel entity.
type Channel struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the Channel entity
// @receiver Channel
// @return []ent.Mixin
func (Channel) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the Channel.
func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("comment").Optional(),
	}
}
