package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/tikafog/of/dbc/mixin"
)

// Update holds the schema definition for the Update entity.
type Update struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the Update entity
// @receiver Update
// @return []ent.Mixin
func (Update) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the Update.
func (Update) Fields() []ent.Field {
	return []ent.Field{
		field.String("os").Default(""),
		field.String("arch").Default(""),
		field.String("version").Default(""),
		field.String("rid").Default(""),
		field.String("crc32").Default(""),
		field.String("attr").Default("core"), //core,app
		field.Bool("forcibly").Default(false),
		field.Bool("truncate").Default(false),
		field.String("title").Default(""),
		field.String("detail").Default(""),
		field.String("sign").Optional(),
	}
}
