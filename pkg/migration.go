package pkg

import (
	"app/ent"
	"context"
	"log"
	"os"

	atlas "ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/lib/pq"
)

func Migration(ctx context.Context, client *ent.Client) {
	dir, err := atlas.NewLocalDir("migrations")
	if err != nil {
		log.Fatal("Failed creating atlas migration directory: ", err)
	}

	var name = os.Getenv("NAME")
	if name == "" {
		if name = os.Getenv("MIGRATION_NAME"); name == "" {
			log.Fatal("Failed generating migration: Missing name")
		}
	}

	if err := client.Debug().Schema.NamedDiff(
		ctx,
		name,
		schema.WithDir(dir),
		schema.WithMigrationMode(schema.ModeReplay),
		schema.WithDialect(dialect.Postgres),
		schema.WithFormatter(atlas.DefaultFormatter),
	); err != nil {
		log.Fatal("Failed running schema migration: ", err)
	}
}
