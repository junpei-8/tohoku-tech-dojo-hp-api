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

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			Annotations(
				entgql.OrderField("NAME"),
			),
		field.String("email").
			Unique().
			NotEmpty().
			Immutable().
			Annotations(
				entgql.OrderField("EMAIL"),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Time("updated_at").
			UpdateDefault(time.Now).
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("UPDATED_AT"),
			),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("created_news_list", News.Type).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.
			Fields("email"),
	}
}

func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
