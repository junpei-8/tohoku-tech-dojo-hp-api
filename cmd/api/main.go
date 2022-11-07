package main

import (
	"app/cmd"
	"app/ent"
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "app/ent/runtime"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/proxy"
	goauth "golang.org/x/oauth2/google"
)

func main() {
	cmd.LoadEnv()

	const SQLScope = "https://www.googleapis.com/auth/sqlservice.admin"

	ctx := context.Background()

	authClient, err := goauth.DefaultClient(ctx, SQLScope)
	if err != nil {
		log.Fatal("connecting google auth", err)
	}

	proxy.Init(authClient, nil, nil)

	dsn := fmt.Sprintf("host=%s:%s:%s user=%s dbname=%s password=%s sslmode=disable",
		"yourGCPProjectID", "asia-northeast1", "yourinstancename", "dbusername", "dbname", "dbpassword",
	)

	db, err := sql.Open("cloudsqlpostgres", dsn)
	if err != nil {
		log.Fatal("opening connection to cloud sql with postgres", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	cmd.Migration(client, ctx)

	cmd.SetupServer(client)
}
