package cmd

import (
	"app/ent"
	"context"
	"log"

	"entgo.io/ent/dialect/sql/schema"
)

func Migration(client *ent.Client, context context.Context) {
	opts := []schema.MigrateOption{
		schema.WithDropColumn(true),
		schema.WithDropIndex(true),
		schema.WithGlobalUniqueID(false),
	}

	if err := client.Schema.Create(
		context,
		opts...,
	); err != nil {
		log.Fatal("running schema migration", err)
	}
}
