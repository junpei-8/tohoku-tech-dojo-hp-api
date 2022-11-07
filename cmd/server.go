package cmd

import (
	"app/ent"
	"app/ent/hook"
	"app/resolvers"
	"net/http"
	"os"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
)

func SetupServer(client *ent.Client) {
	schema := resolvers.NewSchema(client)

	srv := handler.NewDefaultServer(schema)

	srv.Use(entgql.Transactioner{TxOpener: client})

	if os.Getenv("DEBUG") != "false" {
		srv.Use(&debug.Tracer{})
	}

	client.News.Use(hook.News.Hooks(hook.News{})...)

	http.Handle("/",
		playground.Handler("GraphQL playground", "/query"),
	)

	http.Handle("/query", srv)
}
