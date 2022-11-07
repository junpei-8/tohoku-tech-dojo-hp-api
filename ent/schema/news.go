package schema

import (
	"math/rand"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/index"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type News struct {
	ent.Schema
}

func (News) Fields() []ent.Field {
	return []ent.Field{
		field.Int("uid").
			DefaultFunc(func() int {
				rand.Seed(time.Now().UnixNano())
				return rand.Intn(89999999) + 10000000
			}).
			Unique().
			Immutable().
			Annotations(
				entgql.OrderField("UID"),
			),

		field.String("title").
			NotEmpty().
			Annotations(
				entgql.OrderField("TITLE"),
			),

		field.Text("html").
			NotEmpty().
			Annotations(
				entgql.OrderField("HTML"),
			),

		field.Text("markdown").
			NotEmpty().
			Annotations(
				entgql.OrderField("MARKDOWN"),
			),

		field.Int("creator_id").
			Immutable().
			Annotations(
				entgql.OrderField("CREATOR_ID"),
			),

		// 複数持つ必要があるため、別テーブルの実装が必要
		// field.Int("updater_id").
		// 	Optional().
		// 	Nillable().
		// 	Annotations(
		// 		entgql.OrderField("UPDATER_ID"),
		// 	),

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

func (News) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("created_news_list").
			Field("creator_id").
			Unique().
			Required().
			Immutable().
			Annotations(entgql.MapsTo("newsCreator")),

		edge.To("tags", Tag.Type).
			Through("news_tagging", NewsTagging.Type).
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				},
			),
	}
}

func (News) Indexes() []ent.Index {
	return []ent.Index{
		index.
			Fields("uid", "title", "creator_id").
			Edges("user"),
	}
}

func (News) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "news_list"},
		entgql.QueryField("news_list"),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
