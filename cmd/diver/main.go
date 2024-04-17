package main

import (
	"context"

	"github.com/imirjar/Michman/internal/diver/app"
)

func main() {
	if err := app.NewApp().Run(context.Background()); err != nil {
		panic(err.Error())
	}
}
