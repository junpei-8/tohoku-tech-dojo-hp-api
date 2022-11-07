package cmd

import (
	"app/ent"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgres() *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("opening connection to postgres", err)
	}

	return client
}
