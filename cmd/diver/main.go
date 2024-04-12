package main

import (
	"context"

	"github.com/imirjar/Michman/internal/diver/app"
)

func main() {
	app := app.NewApp("localhost:8080/")
	app.Run(context.Background())
}
