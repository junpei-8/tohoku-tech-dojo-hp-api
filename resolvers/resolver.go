package resolvers

import (
	"app"
	"app/ent"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return app.NewExecutableSchema(app.Config{
		Resolvers: &Resolver{client},
	})
}
