package diver

import (
	"github.com/caarlos0/env/v10"
)

type DiverConfig struct {
	Port     string `env:"PORT"`
	Secret   string `env:"DB_PATH"`
	TargetDB string `env:"SECRET"`
	Michman  string `env:"MICHMAN_ADDR"`
}

func NewDiverConfig() *DiverConfig {
	conf := DiverConfig{}
	if err := env.Parse(&conf); err != nil {
		panic(err)
	}
	return &conf
}
