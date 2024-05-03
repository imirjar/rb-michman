package main

import (
	"context"

	"github.com/imirjar/Michman/internal/michman/app"
)

func main() {
	if err := app.NewApp().Run(context.Background()); err != nil {
		panic(err.Error())
	}
}
