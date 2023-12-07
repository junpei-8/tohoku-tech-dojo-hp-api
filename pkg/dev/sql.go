package dev

import (
	"app/ent"
	"log"
	"os"
)

func ConnectPostgres() *ent.Client {
	client, err := ent.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		log.Fatal("Failed opening connection to postgres: ", err)
	}

	return client
}
