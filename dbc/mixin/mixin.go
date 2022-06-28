package mixin

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/google/uuid"
)

type uuidMixin struct {
	mixin.Schema
}

func (uuidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).Unique(),
	}
}

type timeMixin struct {
	mixin.Schema
}

func (timeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("created_unix").Default(0),
		field.Int64("updated_unix").Default(0),
		field.Int64("deleted_unix").Optional(),
	}
}

func BasicSchemaMixin() []ent.Mixin {
	return []ent.Mixin{
		uuidMixin{},
		timeMixin{},
	}
}
