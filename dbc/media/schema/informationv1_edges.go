package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the InformationV1.
func (InformationV1) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("top_lists", TopList.Type),
		edge.From("channel", Channel.Type).Field("channel_id").Ref("informations").Unique(),
	}
}
