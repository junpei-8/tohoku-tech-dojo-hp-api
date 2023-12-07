package hook

import (
	"app/ent"
	"app/ent/user"
	"context"
)

func UserHooks(client *ent.Client) []ent.Hook {
	return []ent.Hook{
		On(
			func(next ent.Mutator) ent.Mutator {
				return UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
					email, _ := m.Email()

					existsUser, err := client.
						User.
						Query().
						Where(user.Email(email)).
						Limit(1).
						All(ctx)

					if err != nil {
						return nil, err
					}

					if len(existsUser) != 0 {
						return existsUser[0], err
					}

					return next.Mutate(ctx, m)
				})
			},

			ent.OpCreate,
		),
	}
}
