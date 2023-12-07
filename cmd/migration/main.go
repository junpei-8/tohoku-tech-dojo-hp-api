package main

import (
	"app/pkg"
	"app/pkg/dev"
	"context"
)

func main() {
	pkg.LoadEnv()

	ctx := context.Background()

	client := dev.ConnectPostgres()

	pkg.Migration(ctx, client)

	defer client.Close()
}
