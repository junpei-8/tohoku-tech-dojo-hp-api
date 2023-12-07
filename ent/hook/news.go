package hook

import (
	"app/ent"
	"app/ent/tag"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

func NewsHooks(client *ent.Client) []ent.Hook {
	return []ent.Hook{
		On(
			func(next ent.Mutator) ent.Mutator {
				return NewsFunc(func(ctx context.Context, m *ent.NewsMutation) (ent.Value, error) {
					// 型アサーション
					// intTagIds := m.TagsIDs()
					// tagIds := make([]any, len(intTagIds))
					tags, err :=
						client.
							Tag.
							Query().
							Where(
								tag.HasNewsTaggingWith(
									func(s *sql.Selector) {
										// s.Where(
										// 	s.
										// 		Select(newstagging.TagFieldID).
										// 		GroupBy(newstagging.TagFieldID).
										// 		Having(
										// 			sql.LT(sql.Count(newstagging.TagFieldID), 2),
										// 		),
										// ),
									},
								),
							).
							All(ctx)

					if err != nil {
						return nil, err
					}

					for i, tag := range tags {
						fmt.Printf("\ntag %v: %v \n", i, tag)
					}

					// return next.Mutate(ctx, m)
					return nil, nil
				})
			},

			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
