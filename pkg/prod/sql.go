package prod

import (
	"app/ent"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"cloud.google.com/go/cloudsqlconn"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func ConnectPostgres() *ent.Client {
	var (
		dbName      = os.Getenv("DB_NAME")          // e.g. 'postgres'
		sqlUser     = os.Getenv("GCP_SQL_USER")     // e.g. 'my-db-user'
		sqlPassword = os.Getenv("GCP_SQL_PASSWORD") // e.g. 'my-db-password'
		sqlInstance = os.Getenv("GCP_SQL_INSTANCE") // e.g. 'project:region:instance'
	)

	dsn := fmt.Sprintf("database=%s user=%s password=%s", dbName, sqlUser, sqlPassword)
	config, err := pgx.ParseConfig(dsn)
	if err != nil {
		log.Fatal(err)
	}

	config.DialFunc = func(ctx context.Context, network, instance string) (net.Conn, error) {
		d, err := cloudsqlconn.NewDialer(ctx)
		if err != nil {
			return nil, err
		}
		return d.Dial(ctx, sqlInstance)
	}

	dbURI := stdlib.RegisterConnConfig(config)
	db, err := sql.Open("pgx", dbURI)
	if err != nil {
		log.Fatal("Failed postgres connection: ", err)
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	client := ent.NewClient(ent.Driver(drv))

	return client
}
