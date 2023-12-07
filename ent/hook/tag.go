package hook

import (
	"app/ent"
	"app/ent/tag"
	"context"
)

func TagHooks(client *ent.Client) []ent.Hook {
	return []ent.Hook{
		On(
			func(next ent.Mutator) ent.Mutator {
				return TagFunc(func(ctx context.Context, m *ent.TagMutation) (ent.Value, error) {
					title, _ := m.Title()

					existsTags, err := client.
						Tag.
						Query().
						Where(tag.Title(title)).
						Limit(1).
						All(ctx)

					if err != nil {
						return nil, err
					}

					if len(existsTags) != 0 {
						return existsTags[0], nil
					}

					return next.Mutate(ctx, m)
				})
			},

			ent.OpCreate,
		),
	}
}
