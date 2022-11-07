// Code generated by ent, DO NOT EDIT.

package newstagging

import (
	"app/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// NewsID applies equality check predicate on the "news_id" field. It's identical to NewsIDEQ.
func NewsID(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNewsID), v))
	})
}

// TagID applies equality check predicate on the "tag_id" field. It's identical to TagIDEQ.
func TagID(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// NewsIDEQ applies the EQ predicate on the "news_id" field.
func NewsIDEQ(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNewsID), v))
	})
}

// NewsIDNEQ applies the NEQ predicate on the "news_id" field.
func NewsIDNEQ(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNewsID), v))
	})
}

// NewsIDIn applies the In predicate on the "news_id" field.
func NewsIDIn(vs ...int) predicate.NewsTagging {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldNewsID), v...))
	})
}

// NewsIDNotIn applies the NotIn predicate on the "news_id" field.
func NewsIDNotIn(vs ...int) predicate.NewsTagging {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldNewsID), v...))
	})
}

// TagIDEQ applies the EQ predicate on the "tag_id" field.
func TagIDEQ(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// TagIDNEQ applies the NEQ predicate on the "tag_id" field.
func TagIDNEQ(v int) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagID), v))
	})
}

// TagIDIn applies the In predicate on the "tag_id" field.
func TagIDIn(vs ...int) predicate.NewsTagging {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTagID), v...))
	})
}

// TagIDNotIn applies the NotIn predicate on the "tag_id" field.
func TagIDNotIn(vs ...int) predicate.NewsTagging {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.NewsTagging(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTagID), v...))
	})
}

// HasNews applies the HasEdge predicate on the "news" edge.
func HasNews() predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, NewsColumn),
			sqlgraph.To(NewsInverseTable, NewsFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NewsTable, NewsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNewsWith applies the HasEdge predicate on the "news" edge with a given conditions (other predicates).
func HasNewsWith(preds ...predicate.News) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, NewsColumn),
			sqlgraph.To(NewsInverseTable, NewsFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, NewsTable, NewsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TagColumn),
			sqlgraph.To(TagInverseTable, TagFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, TagColumn),
			sqlgraph.To(TagInverseTable, TagFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.NewsTagging) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.NewsTagging) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
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
func Not(p predicate.NewsTagging) predicate.NewsTagging {
	return predicate.NewsTagging(func(s *sql.Selector) {
		p(s.Not())
	})
}