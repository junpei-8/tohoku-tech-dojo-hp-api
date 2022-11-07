package hook

import (
	"app/ent"
	"app/ent/schema"
	"context"
	"fmt"
)

// @see https://stackoverflow.com/questions/28800672/how-to-add-new-methods-to-an-existing-type-in-go
type News struct {
	*schema.News
}

func (News) Hooks() []ent.Hook {
	return []ent.Hook{
		On(
			func(next ent.Mutator) ent.Mutator {
				return NewsFunc(func(ctx context.Context, m *ent.NewsMutation) (ent.Value, error) {
					fmt.Printf("Hello from News hook!")
					return next.Mutate(ctx, m)
				})
			},

			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}
