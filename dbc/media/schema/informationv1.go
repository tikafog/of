package schema

import (
	"fmt"
	"image"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"

	"github.com/tikafog/of/dbc/mixin"
)

const (
	PermissionFree = "free"
	PermissionVIP  = "vip"
)

// SkipFrame ...
// @Description: skip to the next frame
type SkipFrame struct {
	Index   int         `json:"index,omitempty"`
	Time    int         `json:"time,omitempty"`
	Start   image.Point `json:"start,omitempty"`
	End     image.Point `json:"end,omitempty"`
	Caption string      `json:"caption,omitempty"`
}

func (s SkipFrame) String() string {
	return fmt.Sprintf("SkipFrame(Index:%d,Time:%d,Start:%v,End:%v,Caption:%v)",
		s.Index,
		s.Time,
		s.Start,
		s.End,
		s.Caption)
}

// InformationV1 holds the schema definition for the InformationV1 entity.
type InformationV1 struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the InformationV1 entity
// @receiver InformationV1
// @return []ent.Mixin
func (InformationV1) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the InformationV1.
func (InformationV1) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("channel_id", uuid.UUID{}).Optional(),
		field.String("root").Default(""),
		field.String("thumb").Default(""),
		field.String("thumb_path").Default(""),
		field.String("poster").Default(""),
		field.String("poster_path").Default(""),
		field.String("media").Default(""),
		field.String("media_path").Default(""),
		field.String("media_index").Default(""),
		field.String("frames").Default(""),
		field.String("frames_path").Default(""),
		field.JSON("frames_particulars", []SkipFrame{}).Optional(),
		field.String("title").Default(""),
		field.String("video_no").Default(""),
		field.Text("intro").Default(""),
		field.Strings("alias"),
		field.Strings("role"),
		field.String("director").Default(""),
		field.String("systematics").Default(""),
		field.String("producer").Default(""),
		field.String("publisher").Default(""),
		field.String("sort_type").Default(""),
		field.String("caption").Default(""),
		field.String("group").Default(""),
		field.String("index").Default(""),
		field.String("release_date").Default(""),
		field.String("format").Default(""),
		field.String("series").Default(""),
		field.Strings("tags"),
		field.String("length").Default(""),
		field.Strings("sample"),
		field.String("uncensored").Default(""),
		field.String("season").Default(""),
		field.String("total_episode").Default(""),
		field.String("episode").Default(""),
		field.String("language").Default(""),
		field.String("sharpness").Default(""),
		field.Bool("watermark").Default(false),
		field.Enum("permission").Values(
			PermissionFree,
			PermissionVIP,
		).Default(PermissionFree),
		field.String("sign").Optional(),      //information check sign
		field.Int("total_blocks").Default(0), //the max block count
	}
}

func (InformationV1) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("channel_id"),
		index.Fields("root"),
	}
}
