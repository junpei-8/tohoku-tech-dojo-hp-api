package main

import (
	"app/pkg"
	"app/pkg/prod"
	"net/http"
	"os"

	"github.com/rs/cors"
)

func main() {
	pkg.LoadEnv()

	client := prod.ConnectPostgres()

	defer client.Close()

	mux := http.NewServeMux()

	prod.SetupGraphqlHandler(client, mux)

	pkg.SetupHandler(mux)

	pkg.SetupHooks(client)

	port := os.Getenv("GCP_API_PORT")

	cors := cors.Default();
	http.ListenAndServe(":"+port, cors.Handler(mux))
}
