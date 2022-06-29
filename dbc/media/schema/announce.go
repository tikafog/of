package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"

	"github.com/tikafog/of/dbc/mixin"
)

const (
	AnnounceKindNotice       = "notice"       //系统通知
	AnnounceKindEvent        = "event"        //折扣
	AnnounceKindAnnouncement = "announcement" //公告
	AnnounceKindUpdate       = "update"       //更新
)

// Announce holds the schema definition for the Announce entity.
type Announce struct {
	ent.Schema
}

// Mixin ...
// @Description: Mixin for the Announce entity
// @receiver Announce
// @return []ent.Mixin
func (Announce) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the Announce.
func (Announce) Fields() []ent.Field {
	return []ent.Field{
		field.String("announce_no").Default(""),
		field.Text("title").Default(""),
		field.Enum("kind").Values(
			AnnounceKindNotice, AnnounceKindEvent, AnnounceKindAnnouncement, AnnounceKindUpdate,
		).Default(AnnounceKindNotice),
		field.Text("content").Default(""),
		field.String("link").Default(""),
		field.String("sign").Optional(), // check sign
		//field.Bool("visited").Default(false),
	}
}

// Edges of the Announce.
func (Announce) Edges() []ent.Edge {
	return nil
}
