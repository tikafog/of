package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/tikafog/of/dbc/mixin"
)

// Page holds the schema definition for the Page entity.
type Page struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the Page entity
// @receiver Page
// @return []ent.Mixin
func (Page) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the HomePage.
func (Page) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("parent_id", uuid.UUID{}).Optional(),
		field.String("title").Default("home"),
		//field.String("meta_type").Default(""),
		//field.String("content").Default(""),
		//field.String("description").Default(""),
		field.Int8("featured_index").Default(0),
		field.String("featured_content").Default(""),
		field.Int64("recommend").Default(0),
	}
}
