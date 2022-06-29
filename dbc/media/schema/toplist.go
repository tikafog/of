package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"github.com/tikafog/of/dbc/mixin"
)

const (
	CategoryNewest    = "newest"    //最新
	CategoryHot       = "hottest"   //最热
	CategorySpecial   = "special"   //推荐
	CategoryFeatured  = "featured"  //推荐
	CategoryStar      = "star"      //明星
	CategoryProducer  = "producer"  //发行商
	CategoryExclusive = "exclusive" //独家
	CategoryTop       = "top"       //置顶
	CategoryNormal    = "normal"    //普通
)

const (
	CornerFree       = "free"       //限免
	CornerDiscount   = "discount"   //折扣
	CornerEvent      = "event"      //活动
	CornerPremium    = "premium"    //精品
	CornerCollection = "collection" //收藏
	CornerLiked      = "liked"      //喜欢
	CornerNone       = "none"       //无
)

// TopList holds the schema definition for the TopList entity.
type TopList struct {
	ent.Schema
}

func (TopList) Mixin() []ent.Mixin {
	return mixin.BasicSchemaMixin()
}

// Fields of the TopList.
func (TopList) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("information_id", uuid.UUID{}).Optional(),
		field.UUID("page_id", uuid.UUID{}).Optional(),
		field.String("title").Default(""),
		field.String("intro").Default(""),
		field.Enum("category").Values(
			CategoryTop, CategoryNewest, CategoryHot, CategoryStar, CategorySpecial, CategoryFeatured,
			CategoryProducer, CategoryExclusive, CategoryNormal,
		).Default(CategoryNormal),
		field.Enum("lower_banner").Values(
			CornerFree, CornerDiscount, CornerEvent,
			CornerPremium, CornerCollection, CornerLiked, CornerNone).
			Default(CornerNone),
		field.Enum("top_right").Values(
			CornerFree, CornerDiscount, CornerEvent,
			CornerPremium, CornerCollection, CornerLiked, CornerNone).
			Default(CornerNone),
	}
}

// Edges of the TopList.
func (TopList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("information", InformationV1.Type).Field("information_id").Ref("top_lists").Unique(),
		edge.From("page", Page.Type).Field("page_id").Ref("top_lists").Unique(),
	}
}
