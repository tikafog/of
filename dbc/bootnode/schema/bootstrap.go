package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

const (
	LevelCore   = "core"   //can request to admin
	LevelSpeed  = "speed"  //can transfer admin info
	LevelNormal = "normal" //only for connect
)

// Bootstrap holds the schema definition for the Bootstrap entity.
type Bootstrap struct {
	ent.Schema
}

// Fields of the Bootstrap.
func (Bootstrap) Fields() []ent.Field {
	return []ent.Field{
		field.String("pid").Unique(),
		field.Strings("addrs"),
		field.Bool("expired").Default(false),
		field.Enum("level").Values(
			LevelCore, LevelSpeed,
			LevelNormal,
		).Default(LevelCore),
		field.Int("service_port").Default(0),
		field.Int("fail_counts").Default(0),
		field.String("sign").Optional(),
		field.Int64("updated_unix").Default(0),
	}
}

// Edges of the Bootstrap.
func (Bootstrap) Edges() []ent.Edge {
	return nil
}
