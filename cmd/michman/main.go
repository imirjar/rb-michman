package main

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/app"
)

func main() {
	app := app.NewApp("localhost:9090/")
	app.Run(context.Background())
}
