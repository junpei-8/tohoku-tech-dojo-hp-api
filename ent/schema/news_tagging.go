package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type NewsTagging struct {
	ent.Schema
}

func (NewsTagging) Fields() []ent.Field {
	return []ent.Field{
		field.Int("news_id").
			Immutable().
			Annotations(
				entgql.OrderField("NEWS_ID"),
			),
		field.Int("tag_id").
			Immutable().
			Annotations(
				entgql.OrderField("TAG_ID"),
			),
	}
}

func (NewsTagging) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("news", News.Type).
			Field("news_id").
			Unique().
			Required().
			Immutable().
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),

		edge.To("tag", Tag.Type).
			Field("tag_id").
			Unique().
			Required().
			Immutable().
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
	}
}

func (NewsTagging) Indexes() []ent.Index {
	return []ent.Index{
		index.
			Fields("news_id", "tag_id").
			Edges("news", "tag"),
	}
}

func (NewsTagging) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("news_id", "tag_id"),
		entgql.QueryField(),
	}
}
