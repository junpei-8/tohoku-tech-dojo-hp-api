package hook

import (
	"app/ent"
	"app/ent/tag"
	"context"
)

type Tag struct {
	*ent.Client
}

func (s Tag) Hooks() []ent.Hook {
	return []ent.Hook{
		On(
			func(next ent.Mutator) ent.Mutator {
				return TagFunc(func(ctx context.Context, m *ent.TagMutation) (ent.Value, error) {
					title, _ := m.Title()

					existsTag, err := s.
						Tag.
						Query().
						Where(tag.Title(title)).
						Limit(1).
						All(ctx)

					if err != nil {
						return nil, err
					}

					if len(existsTag) != 0 {
						return existsTag[0], err
					}

					return next.Mutate(ctx, m)
				})
			},

			ent.OpCreate,
		),
	}
}
