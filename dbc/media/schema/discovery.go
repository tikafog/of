package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/tikafog/of/dbc/mixin"
)

const (
	MediaTypeNone  = "none"
	MediaTypeVideo = "video"
	MediaTypePhoto = "photo"
	MediaTypeBoth  = "both"
)

// Discovery holds the schema definition for the Discovery entity.
type Discovery struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the Discovery entity
// @receiver Discovery
// @return []ent.Mixin
func (Discovery) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the Discovery.
func (Discovery) Fields() []ent.Field {
	return []ent.Field{
		field.String("date").Default(time.RFC3339),
		field.String("rid").Default(""),
		field.Text("title").Default(""),
		field.Text("detail").Default(""),
		field.Enum("mtype").Values(
			MediaTypeNone, MediaTypeVideo, MediaTypePhoto, MediaTypeBoth,
		).Default(MediaTypeBoth),
		field.Strings("links"), //default video/photo,multiple links (0:video,1:photo)
	}
}

// Edges of the Discovery.
func (Discovery) Edges() []ent.Edge {
	return nil
}
