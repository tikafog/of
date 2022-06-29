package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// Edges of the Channel.
func (Channel) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("informations", InformationV1.Type), //.StorageKey(edge.Column("channel_id")),
	}
}
