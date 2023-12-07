package dev

import (
	"app/ent"
	"app/resolver"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/playground"
)

func SetupGraphqlHandler(client *ent.Client, mux *http.ServeMux) {
	schema := resolver.NewSchema(client)

	srv := handler.NewDefaultServer(schema)

	srv.Use(&debug.Tracer{})

	srv.Use(entgql.Transactioner{TxOpener: client})

	mux.Handle("/",
		playground.Handler("GraphQL playground", "/query"),
	)

	mux.Handle("/query", srv)
}
