package main

import (
	"context"

	"github.com/imirjar/Michman/config"
	"github.com/imirjar/Michman/internal/michman/app"
)

func main() {
	cfg := config.NewMichmanConfig()
	app := app.NewApp(cfg)
	app.Run(context.Background())
}
