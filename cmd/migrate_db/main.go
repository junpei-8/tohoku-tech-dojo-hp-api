package main

import (
	"app/cmd"
	"context"
)

func main() {
	cmd.LoadEnv()

	client := cmd.ConnectPostgres()

	ctx := context.Background()

	cmd.Migration(client, ctx)
}
