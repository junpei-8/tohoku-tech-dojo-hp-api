package main

import (
	"app/pkg"
	"app/pkg/dev"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	pkg.LoadEnv()

	client := dev.ConnectPostgres()

	defer client.Close()

	mux := http.NewServeMux()

	dev.SetupGraphqlHandler(client, mux)

	pkg.SetupHandler(mux)

	pkg.SetupHooks(client)

	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")

	log.Printf("listening on http://%s:%s", host, port)

	cors := cors.Default();
	if err := http.ListenAndServe(":"+port, cors.Handler(mux)); err != nil {
		log.Fatal("Terminated http server: ", err)
	}
}
