// Code generated by ent, DO NOT EDIT.

package news

import (
	"app/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// UID applies equality check predicate on the "uid" field. It's identical to UIDEQ.
func UID(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// HTML applies equality check predicate on the "html" field. It's identical to HTMLEQ.
func HTML(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTML), v))
	})
}

// Markdown applies equality check predicate on the "markdown" field. It's identical to MarkdownEQ.
func Markdown(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkdown), v))
	})
}

// CreatorID applies equality check predicate on the "creator_id" field. It's identical to CreatorIDEQ.
func CreatorID(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorID), v))
	})
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UIDEQ applies the EQ predicate on the "uid" field.
func UIDEQ(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUID), v))
	})
}

// UIDNEQ applies the NEQ predicate on the "uid" field.
func UIDNEQ(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUID), v))
	})
}

// UIDIn applies the In predicate on the "uid" field.
func UIDIn(vs ...int) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUID), v...))
	})
}

// UIDNotIn applies the NotIn predicate on the "uid" field.
func UIDNotIn(vs ...int) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUID), v...))
	})
}

// UIDGT applies the GT predicate on the "uid" field.
func UIDGT(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUID), v))
	})
}

// UIDGTE applies the GTE predicate on the "uid" field.
func UIDGTE(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUID), v))
	})
}

// UIDLT applies the LT predicate on the "uid" field.
func UIDLT(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUID), v))
	})
}

// UIDLTE applies the LTE predicate on the "uid" field.
func UIDLTE(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUID), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// HTMLEQ applies the EQ predicate on the "html" field.
func HTMLEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHTML), v))
	})
}

// HTMLNEQ applies the NEQ predicate on the "html" field.
func HTMLNEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHTML), v))
	})
}

// HTMLIn applies the In predicate on the "html" field.
func HTMLIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldHTML), v...))
	})
}

// HTMLNotIn applies the NotIn predicate on the "html" field.
func HTMLNotIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldHTML), v...))
	})
}

// HTMLGT applies the GT predicate on the "html" field.
func HTMLGT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHTML), v))
	})
}

// HTMLGTE applies the GTE predicate on the "html" field.
func HTMLGTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHTML), v))
	})
}

// HTMLLT applies the LT predicate on the "html" field.
func HTMLLT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHTML), v))
	})
}

// HTMLLTE applies the LTE predicate on the "html" field.
func HTMLLTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHTML), v))
	})
}

// HTMLContains applies the Contains predicate on the "html" field.
func HTMLContains(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHTML), v))
	})
}

// HTMLHasPrefix applies the HasPrefix predicate on the "html" field.
func HTMLHasPrefix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHTML), v))
	})
}

// HTMLHasSuffix applies the HasSuffix predicate on the "html" field.
func HTMLHasSuffix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHTML), v))
	})
}

// HTMLEqualFold applies the EqualFold predicate on the "html" field.
func HTMLEqualFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHTML), v))
	})
}

// HTMLContainsFold applies the ContainsFold predicate on the "html" field.
func HTMLContainsFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHTML), v))
	})
}

// MarkdownEQ applies the EQ predicate on the "markdown" field.
func MarkdownEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMarkdown), v))
	})
}

// MarkdownNEQ applies the NEQ predicate on the "markdown" field.
func MarkdownNEQ(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMarkdown), v))
	})
}

// MarkdownIn applies the In predicate on the "markdown" field.
func MarkdownIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldMarkdown), v...))
	})
}

// MarkdownNotIn applies the NotIn predicate on the "markdown" field.
func MarkdownNotIn(vs ...string) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldMarkdown), v...))
	})
}

// MarkdownGT applies the GT predicate on the "markdown" field.
func MarkdownGT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMarkdown), v))
	})
}

// MarkdownGTE applies the GTE predicate on the "markdown" field.
func MarkdownGTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMarkdown), v))
	})
}

// MarkdownLT applies the LT predicate on the "markdown" field.
func MarkdownLT(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMarkdown), v))
	})
}

// MarkdownLTE applies the LTE predicate on the "markdown" field.
func MarkdownLTE(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMarkdown), v))
	})
}

// MarkdownContains applies the Contains predicate on the "markdown" field.
func MarkdownContains(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMarkdown), v))
	})
}

// MarkdownHasPrefix applies the HasPrefix predicate on the "markdown" field.
func MarkdownHasPrefix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMarkdown), v))
	})
}

// MarkdownHasSuffix applies the HasSuffix predicate on the "markdown" field.
func MarkdownHasSuffix(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMarkdown), v))
	})
}

// MarkdownEqualFold applies the EqualFold predicate on the "markdown" field.
func MarkdownEqualFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMarkdown), v))
	})
}

// MarkdownContainsFold applies the ContainsFold predicate on the "markdown" field.
func MarkdownContainsFold(v string) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMarkdown), v))
	})
}

// CreatorIDEQ applies the EQ predicate on the "creator_id" field.
func CreatorIDEQ(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatorID), v))
	})
}

// CreatorIDNEQ applies the NEQ predicate on the "creator_id" field.
func CreatorIDNEQ(v int) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatorID), v))
	})
}

// CreatorIDIn applies the In predicate on the "creator_id" field.
func CreatorIDIn(vs ...int) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatorID), v...))
	})
}

// CreatorIDNotIn applies the NotIn predicate on the "creator_id" field.
func CreatorIDNotIn(vs ...int) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatorID), v...))
	})
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreatedAt), v...))
	})
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreatedAt), v))
	})
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreatedAt), v))
	})
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.News {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUpdatedAt), v...))
	})
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdatedAt), v))
	})
}

// UpdatedAtIsNil applies the IsNil predicate on the "updated_at" field.
func UpdatedAtIsNil() predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUpdatedAt)))
	})
}

// UpdatedAtNotNil applies the NotNil predicate on the "updated_at" field.
func UpdatedAtNotNil() predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUpdatedAt)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTags applies the HasEdge predicate on the "tags" edge.
func HasTags() predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagsWith applies the HasEdge predicate on the "tags" edge with a given conditions (other predicates).
func HasTagsWith(preds ...predicate.Tag) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TagsTable, TagsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNewsTagging applies the HasEdge predicate on the "news_tagging" edge.
func HasNewsTagging() predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NewsTaggingTable, NewsTaggingColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, NewsTaggingTable, NewsTaggingColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNewsTaggingWith applies the HasEdge predicate on the "news_tagging" edge with a given conditions (other predicates).
func HasNewsTaggingWith(preds ...predicate.NewsTagging) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NewsTaggingInverseTable, NewsTaggingColumn),
			sqlgraph.Edge(sqlgraph.O2M, true, NewsTaggingTable, NewsTaggingColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.News) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.News) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.News) predicate.News {
	return predicate.News(func(s *sql.Selector) {
		p(s.Not())
	})
}