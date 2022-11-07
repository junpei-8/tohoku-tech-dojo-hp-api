package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Tag struct {
	ent.Schema
}

func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Unique().
			NotEmpty().
			Immutable().
			Annotations(
				entgql.OrderField("TITLE"),
			),
		field.Time("created_at").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("news_list", News.Type).
			Required().
			Ref("tags").
			Through("news_tagging", NewsTagging.Type).
			Annotations(
				entgql.MapsTo("newsTags"),
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
	}
}

func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.
			Fields("id", "title"),
	}
}

func (Tag) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
