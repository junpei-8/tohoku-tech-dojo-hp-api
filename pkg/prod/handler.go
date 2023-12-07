package prod

import (
	"app/ent"
	"app/resolver"
	"net/http"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql/handler"
)

func SetupGraphqlHandler(client *ent.Client, mux *http.ServeMux) {
	schema := resolver.NewSchema(client)

	srv := handler.NewDefaultServer(schema)

	srv.Use(entgql.Transactioner{TxOpener: client})

	mux.Handle("/query", srv)
}
