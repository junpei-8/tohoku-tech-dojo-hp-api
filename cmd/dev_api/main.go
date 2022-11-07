package main

import (
	"app/cmd"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {
	cmd.LoadEnv()

	client := cmd.ConnectPostgres()

	ctx := context.Background()

	cmd.Migration(client, ctx)

	cmd.SetupServer(client)

	port := os.Getenv("API_PORT")

	log.Printf("listening on http://localhost:%s", port)

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}
