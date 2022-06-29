package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the HomePage.
func (Page) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subpages", Page.Type).
			From("parent").
			Unique().
			Field("parent_id"),
		edge.To("top_lists", TopList.Type), //.StorageKey(edge.Column("featured_id")),
	}
}
